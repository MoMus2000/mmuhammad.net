package scripts

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func StatusCheck(reciever string, message string) {
	accountSid := "AC7c1d4068211dfa361cfc6be3a3af78a8"
	authToken := "e941f928a7c1c5cd8f3be9a8ae47e6d3"

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
		fmt.Println(err.Error())
	} else {
		reps, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(reps))
	}
}
