package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jwaggs/ofxgo"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

var wg = &sync.WaitGroup{}

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

// default to chase til we get more of the system carved out
func defineServerFlags(f *flag.FlagSet) {
	// short function to force successful secret reads
	readSecret := func(s string) string {
		data, err := ioutil.ReadFile(s)
		if err != nil {
			log.Fatalln("error reading secret at path:", s, "\nsuppressing error output for security reasons.")
		}
		return string(data)
	}


	f.StringVar(&serverURL, "url", "https://ofx.chase.com", "Financial institution's OFX Server URL (see ofxhome.com if you don't know it)")
	f.StringVar(&clientUID, "clientuid", readSecret("/secrets/auth/clientuid.txt"), "Client UID (only required by a few FIs, like Chase)")
	f.StringVar(&username, "username", readSecret("/secrets/auth/username.txt"), "Your username at financial institution")
	f.StringVar(&password, "password", readSecret("/secrets/auth/password.txt"), "Your password at financial institution")
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

func init() {
	defineServerFlags(scanChaseCmd.Flags)
}

var commands = []command{
	scanChaseCmd,
}

var scanChaseCmd = command{
	Name:        "chase",
	Description: "scan then publish chase accounts and transactions",
	Flags:       flag.NewFlagSet("chase", flag.ExitOnError),
	CheckFlags:  checkServerFlags,
	Do:          getAccounts,
}

func usage() {
	fmt.Println(`Piggy Scan.
Usage:
	scan command [arguments]
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
