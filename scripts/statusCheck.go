package scripts

import (
	"encoding/json"
	"fmt"

	"github.com/dongri/phonenumber"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/xuri/excelize/v2"
)

func ValidatePhoneNumbers(excelFile *excelize.File) {
	number := phonenumber.Parse("+1 647-513 0152", "CA")
	fmt.Println(number)
}

func StatusCheck(reciever string, message string) {
	accountSid := "AC7c1d4068211dfa361cfc6be3a3af78a8"
	authToken := "1cf960e93c262a49bbb97c7696559f47"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(reciever)
	params.SetFrom("+13862515211")
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		// fmt.Println(err.Error())
	} else {
		json.Marshal(*resp)
		// fmt.Println("Response: " + string(response))
	}
}
