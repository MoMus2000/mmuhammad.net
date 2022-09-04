package controllers

import (
	"fmt"
	"mustafa_m/views"
	"net/http"
	"os"
)

func InternalServerError() *views.View {
	return NewStaticController().InternalServerError
}

func ForbiddenError() *views.View {
	return NewStaticController().ForbiddenError
}

func GetIP(r *http.Request) {
	ip := r.RemoteAddr
	xforward := r.Header.Get("X-Forwarded-For")
	ipAddr := fmt.Sprintf("IP: %s", ip)
	forwardFor := fmt.Sprintf("X-Forwarded-For : %s", xforward)
	file, err := os.OpenFile("visitors.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening the file")
	}
	_, err = file.Write([]byte(ipAddr + "\n" + forwardFor + "\n"))
	if err != nil {
		fmt.Println("Error writing to the file")
	}
}

func WrapIPHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call any pre handler functions here
		go GetIP(r)
		f(w, r)
	}
}
