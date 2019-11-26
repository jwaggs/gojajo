package main

import (
	"fmt"
	"github.com/jwaggs/ofxgo"
	"os"
	"time"
)

type bankAcctOpts struct {
	accType string
	accID string
	bnkID string
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
				fmt.Printf("Bank Account:\n\tBankName: %s\n\tBankDesc: %sn\tBankID: %s\n\tAcctID: %s\n\tAcctType: %s\n", acct.Name, acct.Desc, acct.BankAcctInfo.BankAcctFrom.BankID, acct.BankAcctInfo.BankAcctFrom.AcctID, acct.BankAcctInfo.BankAcctFrom.AcctType)
				acct := acct // capture variable
				wg.Add(1)
				go func() {
					defer wg.Done()
					bankTransactions(bankAcctOpts{
						accType: acct.BankAcctInfo.BankAcctFrom.AcctType.String(),
						accID:   acct.BankAcctInfo.BankAcctFrom.AcctID.String(),
						bnkID:   acct.BankAcctInfo.BankAcctFrom.BankID.String(),
					})
				}()

			} else if acct.CCAcctInfo != nil {
				fmt.Printf("Credit card:\n\tAcctID: \"%s\"\n", acct.CCAcctInfo.CCAcctFrom.AcctID)
			} else if acct.InvAcctInfo != nil {
				fmt.Printf("Investment account:\n\tBrokerID: \"%s\"\n\tAcctID: \"%s\"\n", acct.InvAcctInfo.InvAcctFrom.BrokerID, acct.InvAcctInfo.InvAcctFrom.AcctID)
			} else {
				fmt.Printf("Unknown type: %s %s\n", acct.Name, acct.Desc)
			}
		}
	}
	fmt.Println("WAITING")
	wg.Wait()
	fmt.Println("DONE-WAITING")
}