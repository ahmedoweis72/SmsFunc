package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ReceiverMSISDN := 123
	SenderName := "Ahmed"
	SMSText := "it's easy"

	sms(SenderName, ReceiverMSISDN, SMSText)
}

var SenderName string
var ReceiverMSISDN int
var SMSText string

func sms(SenderName string, ReceiverMSISDN int, SMSText string) {
	// Create the request payload as a byte slice
	payload := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<SubmitSMSRequest xmlns="http://www.edafa.com/web2sms/sms/model/">
<AccountId>1</AccountId>
<Password>Password</Password>
<SecureHash></SecureHash>
<SMSList>
 <SenderName>SenderName</SenderName>
 <ReceiverMSISDN>ReceiverMSISDN</ReceiverMSISDN>
 <SMSText>SMSText</SMSText>
</SMSList>
</SubmitSMSRequest>
`)

	// Make the HTTP POST request
	resp, err := http.Post("https://<serverip>:<port>/web2sms/sms/submit/CampaignApi ", "application/xml", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("INVALID_REQUEST:", err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
