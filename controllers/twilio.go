package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mustafa_m/models"
	"mustafa_m/scripts"
	"mustafa_m/views"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Twilio struct {
	LoginPage     *views.View
	ContactUpload *views.View
	FmbService    *models.FmbService
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

func NewTwilioController(FmbService *models.FmbService) *Twilio {
	return &Twilio{
		FmbService:    FmbService,
		LoginPage:     views.NewView("bootstrap", "fmb/login.gohtml"),
		ContactUpload: views.NewView("bootstrap", "fmb/upload.gohtml"),
	}
}

func (tw *Twilio) GetFmbLoginPage(w http.ResponseWriter, r *http.Request) {
	if validateJWTFmb(r) {
		http.Redirect(w, r, "/fmb/upload", http.StatusFound)
		return
	}
	tw.LoginPage.Render(w, nil)
}

func (tw *Twilio) FmbLogin(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	parseForm(r, &form)
	fmt.Println(form)
	fmbTemp := models.Fmb{Email: form.Email, Password: form.Password}
	result, err := tw.FmbService.ByEmail(&fmbTemp)
	if err != nil {
		fmt.Println(err)
		InternalServerError().Render(w, nil)
	}
	fmt.Println(result)

	createJWTFmb(w, &fmbTemp)

	http.Redirect(w, r, "/fmb/upload", http.StatusFound)
}

func (tw *Twilio) GetUploadPage(w http.ResponseWriter, r *http.Request) {
	if !validateJWTFmb(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{FmbLoggedIn: "true"}
	tw.ContactUpload.Render(w, data)
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
	fmt.Println(payload)
	if err != nil {
		InternalServerError().Render(w, nil)
	}
	filename := fmt.Sprintf("./temp/%s_%s_%s.xlsx", payload.SenderName, payload.SenderPhone, uuid.New().String())
	if err := excelFile.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
	// Then send flask request
	flask_payload := make(map[string]string)

	flask_payload["message"] = payload.SenderMessage
	flask_payload["sender"] = payload.SenderPhone
	flask_payload["senderName"] = payload.SenderName
	flask_payload["fileName"] = strings.Split(filename, "/")[2]

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/fmb/send_message", "application/json", bytes.NewBuffer(jsonString))

	if err != nil {
		fmt.Println(err)
	}

	b, err := io.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func AddTwilioRoutes(r *mux.Router, tw *Twilio) {
	// r.HandleFunc("/api/v1/twilio/response", tw.GetWebHookResponse).Methods("GET")
	// r.HandleFunc("/api/v1/twilio/submit", tw.SubmitWebHookResponse).Methods("POST")
	r.HandleFunc("/fmb", tw.GetFmbLoginPage).Methods("GET")
	r.HandleFunc("/fmb", tw.FmbLogin).Methods("POST")
	r.HandleFunc("/fmb/signout", tw.SignoutJWTFmb).Methods("GET")
	r.HandleFunc("/fmb/upload", tw.GetUploadPage).Methods("GET")
	r.HandleFunc("/fmb/upload", tw.UploadContacts).Methods("POST")
	r.HandleFunc("/api/v1/twilio/statusCheck", tw.SampleApiTest).Methods("POST")
}
