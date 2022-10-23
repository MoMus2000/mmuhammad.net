package terminal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"sms.mmuhammad.net/controllers/auth"
)

type SingleSmsPayload struct {
	SenderName    string `json:"SenderName"`
	SenderPhone   string `json:"SenderPhone"`
	SenderMessage string `json:"TextMessage"`
}

type BulkSmsPayload struct {
	SenderName    string `schema:"SenderName"`
	SenderPhone   string `schema:"SenderPhone"`
	SenderMessage string `schema:"TextMessage"`
}

func (sms *SmsTerminal) SendSingleSms(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username

	payload := SingleSmsPayload{}
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

	resp, err := http.Post("http://localhost:3001/api/v1/sms/single_sms", "application/json", bytes.NewBuffer(jsonString))

	if resp.StatusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (sms *SmsTerminal) SendBulkSms(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userEmail := auth.GetEmailFromJwt(r).Username

	payload := BulkSmsPayload{}

	excelFile, err := parseExcelForm(r, &payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filename := fmt.Sprintf("../temp/%s_%s.xlsx", payload.SenderName, payload.SenderPhone)
	if err := excelFile.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
	// Then send flask request
	flask_payload := make(map[string]string)

	flask_payload["message"] = payload.SenderMessage
	flask_payload["sender"] = payload.SenderPhone
	flask_payload["senderName"] = payload.SenderName
	flask_payload["fileName"] = strings.Split(filename, "/")[2]
	flask_payload["email"] = userEmail

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/sms/bulk_sms", "application/json", bytes.NewBuffer(jsonString))

	if resp.StatusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := io.ReadAll(resp.Body)

	fmt.Println(string(b))

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusCreated)
}
