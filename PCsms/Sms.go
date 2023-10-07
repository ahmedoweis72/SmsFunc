package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SMSLists struct {
	XMLName        string `xml:"SMSList"`
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

	funcsms(1, "123@test", "Ahmed", 1, "how are yoy?")
}

func funcsms(Accountid int, Pass string, Name string, Receiver int, Text string) {
	h := sha256.New()
	h.Write([]byte("AccountId=" + string(Accountid) + "&Password=" + Pass + "&SenderName=" + Name + "&ReceiverMSISDN=" + string(Receiver) + "&SMSText=" + Text))
	Secure := hex.EncodeToString(h.Sum(nil))
	fmt.Println(Secure)
	bk := SubmitSMSRequest{
		AccountId:  Accountid,
		Password:   Pass,
		SecureHash: []byte(Secure),
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
	/*resp.RespCode = "00"
	resp.RespMsg = "Login Successful"*/
	status := resp.StatusCode
	if status == 200 {
		fmt.Println("SUCCESS")
	}
	fmt.Println("INVALID_REQUEST")

	// Print the response body

	fmt.Println(string(body))

}
