package scripts

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func StatusCheck(reciever string, message string) error {
	accountSid := os.Getenv("TWILIO_ACCOUNT")
	authToken := os.Getenv("TWILIO_TOKEN")

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
		return err
	} else {
		reps, err := json.Marshal(*resp)
		fmt.Println("Response: " + string(reps))
		return err
	}
}
