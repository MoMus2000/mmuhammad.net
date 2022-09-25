package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"mustafa_m/scripts"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/mux"
)

type Twilio struct {
	ContactUpload *views.View
}

type TwilioPayload struct {
	SenderName    string `json:"SenderName"`
	SenderPhone   string `json:"SenderPhone"`
	SenderMessage string `json:"TextMessage"`
}

type TwilioFilePayload struct {
	SenderName    string `schema:"SenderName"`
	SenderPhone   string `schema:"SenderPhone"`
	SenderMessage string `schema:"TextMessage"`
}

func NewTwilioController() *Twilio {
	return &Twilio{
		ContactUpload: views.NewView("bootstrap", "fmb/upload.gohtml"),
	}
}

func (tw *Twilio) GetUploadPage(w http.ResponseWriter, r *http.Request) {
	tw.ContactUpload.Render(w, nil)
}

func (tw *Twilio) GetWebHookResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}

func (tw *Twilio) SubmitWebHookResponse(w http.ResponseWriter, r *http.Request) {
	bytes, _ := io.ReadAll(r.Body)
	bodyString := string(bytes)
	fmt.Println(bodyString)
}

func (tw *Twilio) SampleApiTest(w http.ResponseWriter, r *http.Request) {
	payload := TwilioPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		InternalServerError().Render(w, nil)
	}
	go scripts.StatusCheck(payload.SenderPhone, payload.SenderMessage)
}

func (tw *Twilio) UploadContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API HIT")
	payload := TwilioFilePayload{}
	excelFile, err := parseExcelForm(r, &payload)
	if err != nil {
		InternalServerError().Render(w, nil)
	}
	rows, err := excelFile.GetRows("Sheet1")
	fmt.Println("Total number of rows : ", len(rows))
}

func AddTwilioRoutes(r *mux.Router, tw *Twilio) {
	// r.HandleFunc("/api/v1/twilio/response", tw.GetWebHookResponse).Methods("GET")
	// r.HandleFunc("/api/v1/twilio/submit", tw.SubmitWebHookResponse).Methods("POST")
	r.HandleFunc("/fmb/upload", tw.GetUploadPage).Methods("GET")
	r.HandleFunc("/fmb/upload", tw.UploadContacts).Methods("POST")
	r.HandleFunc("/api/v1/twilio/statusCheck", tw.SampleApiTest).Methods("POST")
}
