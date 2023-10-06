package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type SMSRequest struct {
	XMLName    xml.Name  `xml:"SMSRequest"`
	AccountId  int       `xml:"AccountId"`
	Password   string    `xml:"Password"`
	SecureHash string    `xml:"SecureHash"`
	SMSLists   []SMSList `xml:"skills"`
}
type SMSList struct {
	XMLName        xml.Name `xml:"SMSList"`
	SenderName     string   `xml:"SenderName"`
	ReceiverMSISDN int      `xml:"ReceiverMSISDN"`
	SMSText        string   `xml:"SMSText"`
}

func main() {
	/*ReceiverMSISDN := 123
	SenderName := "Ahmed"
	SMSText := "it's easy"

	sms(SenderName, ReceiverMSISDN, SMSText)*/
}

/*type smsS struct {
	smsS           string `xml:",innerxml"`
	ReceiverMSISDN string `xml:"ReceiverMSISDN,attr"`
	SenderName     string `xml:"SenderName,attr"`
	SMSText        string `xml:"SMSText,attr"`
}*/

func sms(SenderName string, ReceiverMSISDN int, SMSText string) {
	// Create the request payload as a byte slice
	/*payload := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<SubmitSMSRequest xmlns="http://www.edafa.com/web2sms/sms/model/">
	<AccountId>1</AccountId>
	<Password>Password</Password>
	<SecureHash></SecureHash>
	<SMSList>
	 <SenderName></SenderName>
	 <ReceiverMSISDN></ReceiverMSISDN>
	 <SMSText></SMSText>
	</SMSList>
	</SubmitSMSRequest>
	`)*/
	// Make the HTTP POST request

	smsreq := SMSRequest{
		AccountId:  1,
		Password:   "Password",
		SecureHash: "",
		SMSLists: []SMSList{{
			SenderName:     "",
			ReceiverMSISDN: 1,
			SMSText:        "",
		}},
	}
	xmlFile, err := os.Create("smsreq.xml")
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}
	xmlFile.WriteString(xml.Header)
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&smsreq)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return
	}
	/*resp, err := http.Post("https://<serverip>:<port>/web2sms/sms/submit/CampaignApi ", "smsreq.xml/xml", bytes.NewBuffer(*encoder))
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
	fmt.Println(string(body))*/
}
