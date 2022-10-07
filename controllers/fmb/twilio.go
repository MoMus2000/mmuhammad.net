package fmb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mustafa_m/controllers"
	"mustafa_m/models"
	"mustafa_m/scripts"
	"mustafa_m/views"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var TwilioPhone = "+16476916189"

type Twilio struct {
	LoginPage     *views.View
	ContactUpload *views.View
	FmbDashBoard  *views.View
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
		FmbDashBoard:  views.NewView("bootstrap", "fmb/dashboard.gohtml"),
	}
}

func (tw *Twilio) GetFmbLoginPage(w http.ResponseWriter, r *http.Request) {
	if validateJWTFmb(r) {
		http.Redirect(w, r, "/fmb/upload", http.StatusFound)
		return
	}
	tw.LoginPage.Render(w, nil)
}

func (tw *Twilio) GetFmbDashboard(w http.ResponseWriter, r *http.Request) {
	if !validateJWTFmb(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{FmbLoggedIn: "true"}
	tw.FmbDashBoard.Render(w, data)
}

func (tw *Twilio) FmbLogin(w http.ResponseWriter, r *http.Request) {
	form := controllers.LoginForm{}
	controllers.ParseForm(r, &form)
	fmt.Println(form)
	fmbTemp := models.Fmb{Email: form.Email, Password: form.Password}
	result, err := tw.FmbService.ByEmail(&fmbTemp)
	if err != nil {
		fmt.Println(err)
		controllers.InternalServerError().Render(w, nil)
	}
	fmt.Println(result)

	createJWTFmb(w, &fmbTemp)

	http.Redirect(w, r, "/fmb/upload", http.StatusFound)
}

func (tw *Twilio) GetUploadPage(w http.ResponseWriter, r *http.Request) {
	if !validateJWTFmb(r) {
		controllers.ForbiddenError().Render(w, nil)
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

func (tw *Twilio) GetAvailableBalance(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:3001/api/v1/fmb/app_balance")
	if err != nil {
		fmt.Println(err)
	}
	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	fmt.Println(err)
	b, err := json.Marshal(j)
	fmt.Fprintln(w, string(b))
}

func (tw *Twilio) GetMessageLength(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:3001/api/v1/fmb/get_history")
	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	fmt.Println(err)
	b, err := json.Marshal(j)
	fmt.Fprintln(w, string(b))
}

func (tw *Twilio) SampleApiTest(w http.ResponseWriter, r *http.Request) {
	payload := TwilioPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		controllers.InternalServerError().Render(w, nil)
	}

	err = scripts.StatusCheck(payload.SenderPhone, payload.SenderMessage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (tw *Twilio) UploadContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API HIT")
	payload := TwilioFilePayload{}
	excelFile, err := parseExcelForm(r, &payload)
	fmt.Println(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
	flask_payload["twilioPhone"] = TwilioPhone

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/fmb/send_message", "application/json", bytes.NewBuffer(jsonString))

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

func AddTwilioRoutes(r *mux.Router, tw *Twilio) {
	// r.HandleFunc("/api/v1/twilio/response", tw.GetWebHookResponse).Methods("GET")
	// r.HandleFunc("/api/v1/twilio/submit", tw.SubmitWebHookResponse).Methods("POST")
	r.HandleFunc("/fmb", tw.GetFmbLoginPage).Methods("GET")
	r.HandleFunc("/fmb", tw.FmbLogin).Methods("POST")
	r.HandleFunc("/fmb/signout", tw.SignoutJWTFmb).Methods("GET")
	r.HandleFunc("/fmb/upload", tw.GetUploadPage).Methods("GET")
	r.HandleFunc("/fmb/upload", tw.UploadContacts).Methods("POST")
	r.HandleFunc("/fmb/dashboard", tw.GetFmbDashboard).Methods("GET")
	r.HandleFunc("/api/v1/twilio/statusCheck", tw.SampleApiTest).Methods("POST")
	r.HandleFunc("/api/v1/twilio/balanceCheck", tw.GetAvailableBalance).Methods("GET")
	r.HandleFunc("/api/v1/twilio/getMessageLength", tw.GetMessageLength).Methods("GET")
}
