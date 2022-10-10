package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/controllers/home"
)

func main() {
	r := mux.NewRouter()

	landC := home.NewLandingPageController()

	home.AddHomePageRoutes(r, landC)

	ipAddress := getLocalIpAddress()
	port := "3002"
	fmt.Printf("Listening on %s:%s\n", ipAddress, port)

	http.ListenAndServe(":"+port, r)
}

func getLocalIpAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
