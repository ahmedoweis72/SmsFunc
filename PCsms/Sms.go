package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SMSLists struct {
	XMLName        string `xml:"transition"`
	SenderName     string `xml:"duration"`
	ReceiverMSISDN int    `xml:"ReceiverMSISDN"`
	SMSText        string `xml:"SMSText"`
}

type SubmitSMSRequest struct {
	XMLName    string     `xml:"SubmitSMSRequest"`
	AccountId  int        `xml:"AccountId"`
	Password   string     `xml:"Password"`
	SecureHash []byte     `xml:"SecureHash"`
	SMSList    []SMSLists `xml:"SMSList"`
}

func main() {
	//AccountId=1&Password=password&SenderName=sender1&ReceiverMSISDN=201000000000&
	//SMSText=text1
	Accountid := 1
	Pass := "123@test"
	Name := "Ahmed"
	Receiver := 1
	Text := "how are yoy?"

	h := sha256.New()
	h.Write([]byte(string(Accountid) + Pass + Name + string(Receiver) + Text))
	// Calculate and print the hash
	Secure := h.Sum(nil)
	//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	bk := SubmitSMSRequest{
		AccountId:  Accountid,
		Password:   Pass,
		SecureHash: Secure,
		SMSList: []SMSLists{{
			SenderName:     Name,
			ReceiverMSISDN: Receiver,
			SMSText:        Text,
		}},
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")

	err := enc.Encode(&bk)
	fmt.Println(os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	out, err := xml.Marshal(&bk)
	//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

	resp, err := http.Post("https://<serverip>:<port>/web2sms/sms/submit/CampaignApi ", "smsreq/xml", bytes.NewBuffer(out))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body
	fmt.Println(string(body))
}
