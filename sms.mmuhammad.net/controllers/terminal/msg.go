package terminal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"sms.mmuhammad.net/controllers/auth"
)

type TwilioPayload struct {
	SenderName    string `json:"SenderName"`
	SenderPhone   string `json:"SenderPhone"`
	SenderMessage string `json:"TextMessage"`
}

func (sms *SmsTerminal) SendSingleSms(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username

	payload := TwilioPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	flask_payload := make(map[string]string)

	flask_payload["message"] = payload.SenderMessage
	flask_payload["sender"] = payload.SenderPhone
	flask_payload["senderName"] = payload.SenderName
	flask_payload["email"] = userEmail

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/sms/singleSms", "application/json", bytes.NewBuffer(jsonString))

	fmt.Println(resp.StatusCode)

	w.WriteHeader(http.StatusCreated)
}
