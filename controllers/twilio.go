package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"mustafa_m/models"
	"mustafa_m/scripts"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/mux"
)

type Twilio struct {
	LoginPage     *views.View
	ContactUpload *views.View
	AdminService  *models.AdminService
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

func NewTwilioController(adminService *models.AdminService) *Twilio {
	return &Twilio{
		AdminService:  adminService,
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
	adminTemp := models.Admin{Email: form.Email, Password: form.Password}
	result, err := tw.AdminService.ByEmail(&adminTemp)
	if err != nil {
		fmt.Println(err)
		InternalServerError().Render(w, nil)
	}
	fmt.Println(result)

	createJWTFmb(w, &adminTemp)

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
	if err != nil {
		InternalServerError().Render(w, nil)
	}
	rows, err := excelFile.GetRows("Sheet1")
	fmt.Println("Total number of rows : ", len(rows))
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
