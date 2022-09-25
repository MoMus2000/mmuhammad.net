package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Twilio struct {
}

func (tw *Twilio) GetWebHookResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}

func (tw *Twilio) SubmitWebHookResponse(w http.ResponseWriter, r *http.Request) {
	bytes, _ := io.ReadAll(r.Body)
	bodyString := string(bytes)
	fmt.Println(bodyString)
}

func AddTwilioRoutes(r *mux.Router, tw *Twilio) {
	r.HandleFunc("/api/v1/twilio/response", tw.GetWebHookResponse).Methods("GET")
	r.HandleFunc("/api/v1/twilio/submit", tw.SubmitWebHookResponse).Methods("POST")
}
