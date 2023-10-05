package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	sms()
}

func sms() {
	// Create the request payload as a byte slice
	payload := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<SubmitSMSRequest xmlns="http://www.edafa.com/web2sms/sms/model/">
<AccountId></AccountId>
<Password></Password>
<SecureHash></SecureHash>
<SMSList>
 <SenderName></SenderName>
 <ReceiverMSISDN></ReceiverMSISDN>
 <SMSText></SMSText>
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
