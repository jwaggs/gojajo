package main

import (
	"flag"
	"fmt"
	"github.com/jwaggs/ofxgo"
	"github.com/howeyc/gopass"
	"os"
	"strconv"
	"time"
)

type command struct {
	Name        string
	Description string
	Flags       *flag.FlagSet
	CheckFlags  func() bool // Check the flag values after they're parsed, printing errors and returning false if they're incorrect
	Do          func()      // Run the command (only called if CheckFlags returns true)
}

func (c *command) usage() {
	fmt.Printf("Usage of %s:\n", c.Name)
	c.Flags.PrintDefaults()
}

// flags common to all server transactions
var serverURL, username, password, org, fid, appID, appVer, ofxVersion, clientUID string
var noIndentRequests bool

func defineServerFlags(f *flag.FlagSet) {
	f.StringVar(&serverURL, "url", "https://ofx.chase.com", "Financial institution's OFX Server URL (see ofxhome.com if you don't know it)")
	f.StringVar(&clientUID, "clientuid", os.Getenv("CHASE_CLIENTUID"), "Client UID (only required by a few FIs, like Chase)")
	f.StringVar(&username, "username", os.Getenv("CHASE_USERNAME"), "Your username at financial institution")
	f.StringVar(&password, "password", "", "Your password at financial institution")
	f.StringVar(&org, "org", "B1", "'ORG' for your financial institution")
	f.StringVar(&fid, "fid", "10898", "'FID' for your financial institution")
	f.StringVar(&appID, "appid", "QWIN", "'APPID' to pretend to be")
	f.StringVar(&appVer, "appver", "2700", "'APPVER' to pretend to be")
	f.StringVar(&ofxVersion, "ofxversion", "220", "OFX version to use")
	f.BoolVar(&noIndentRequests, "noindent", false, "Don't indent OFX requests")
}

func checkServerFlags() bool {
	var ret bool = true
	if len(serverURL) == 0 {
		fmt.Println("Error: Server URL empty")
		ret = false
	}
	if len(username) == 0 {
		fmt.Println("Error: Username empty")
		ret = false
	}

	if ret && len(password) == 0 {
		fmt.Printf("Password for %s: ", username)
		pass, err := gopass.GetPasswd()
		if err != nil {
			fmt.Printf("Error reading password: %s\n", err)
			ret = false
		} else {
			password = string(pass)
		}
	}
	return ret
}

func newRequest() (ofxgo.Client, *ofxgo.Request) {
	ver, err := ofxgo.NewOfxVersion(ofxVersion)
	if err != nil {
		fmt.Println("Error creating new OfxVersion enum:", err)
		os.Exit(1)
	}
	var client = ofxgo.GetClient(serverURL,
		&ofxgo.BasicClient{
			AppID:       appID,
			AppVer:      appVer,
			SpecVersion: ver,
			NoIndent:    noIndentRequests,
		})

	var query ofxgo.Request
	query.URL = serverURL
	query.Signon.ClientUID = ofxgo.UID(clientUID)
	query.Signon.UserID = ofxgo.String(username)
	query.Signon.UserPass = ofxgo.String(password)
	query.Signon.Org = ofxgo.String(org)
	query.Signon.Fid = ofxgo.String(fid)

	return client, &query
}

var getAccountsCommand = command{
	Name:        "get-accounts",
	Description: "List accounts at your financial institution",
	Flags:       flag.NewFlagSet("get-accounts", flag.ExitOnError),
	CheckFlags:  checkServerFlags,
	Do:          getAccounts,
}

func init() {
	defineServerFlags(getAccountsCommand.Flags)
}

func getAccounts() {
	client, query := newRequest()

	uid, err := ofxgo.RandomUID()
	if err != nil {
		fmt.Println("Error creating uid for transaction:", err)
		os.Exit(1)
	}

	acctInfo := ofxgo.AcctInfoRequest{
		TrnUID:   *uid,
		DtAcctUp: ofxgo.Date{Time: time.Unix(0, 0)},
	}
	query.Signup = append(query.Signup, &acctInfo)

	response, err := client.Request(query)
	if err != nil {
		fmt.Println("Error requesting account information:", err)
		os.Exit(1)
	}

	if response.Signon.Status.Code != 0 {
		meaning, _ := response.Signon.Status.CodeMeaning()
		fmt.Printf("Nonzero signon status (%d: %s) with message: %s\n", response.Signon.Status.Code, meaning, response.Signon.Status.Message)
		os.Exit(1)
	}

	if len(response.Signup) < 1 {
		fmt.Println("No signup messages received")
		return
	}

	fmt.Printf("\nFound the following accounts:\n\n")

	if acctinfo, ok := response.Signup[0].(*ofxgo.AcctInfoResponse); ok {
		for _, acct := range acctinfo.AcctInfo {
			if acct.BankAcctInfo != nil {
				fmt.Printf("Bank Account:\n\tBankID: \"%s\"\n\tAcctID: \"%s\"\n\tAcctType: %s\n", acct.BankAcctInfo.BankAcctFrom.BankID, acct.BankAcctInfo.BankAcctFrom.AcctID, acct.BankAcctInfo.BankAcctFrom.AcctType)
			} else if acct.CCAcctInfo != nil {
				fmt.Printf("Credit card:\n\tAcctID: \"%s\"\n", acct.CCAcctInfo.CCAcctFrom.AcctID)
			} else if acct.InvAcctInfo != nil {
				fmt.Printf("Investment account:\n\tBrokerID: \"%s\"\n\tAcctID: \"%s\"\n", acct.InvAcctInfo.InvAcctFrom.BrokerID, acct.InvAcctInfo.InvAcctFrom.AcctID)
			} else {
				fmt.Printf("Unknown type: %s %s\n", acct.Name, acct.Desc)
			}
		}
	}
}

var commands = []command{
	getAccountsCommand,
}

func usage() {
	fmt.Println(`The ofxgo command-line client provides a simple interface to
query, parse, and display financial data via the OFX specification.
Usage:
	ofx command [arguments]
The commands are:`)

	maxlen := 0
	for _, cmd := range commands {
		if len(cmd.Name) > maxlen {
			maxlen = len(cmd.Name)
		}
	}
	formatString := "    %-" + strconv.Itoa(maxlen) + "s    %s\n"

	for _, cmd := range commands {
		fmt.Printf(formatString, cmd.Name, cmd.Description)
	}
}

func runCmd(c *command) {
	err := c.Flags.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("Error parsing flags: %s\n", err)
		c.usage()
		os.Exit(1)
	}

	if !c.CheckFlags() {
		fmt.Println()
		c.usage()
		os.Exit(1)
	}

	c.Do()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Error: Please supply a sub-command. Usage:\n\n")
		usage()
		os.Exit(1)
	}
	cmdName := os.Args[1]
	for _, cmd := range commands {
		if cmd.Name == cmdName {
			runCmd(&cmd)
			os.Exit(0)
		}
	}

	switch cmdName {
	case "-h", "-help", "--help", "help":
		usage()
	default:
		fmt.Println("Error: Invalid sub-command. Usage:")
		usage()
		os.Exit(1)
	}
}
