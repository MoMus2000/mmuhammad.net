package main

import (
	"fmt"
	"log"
	"mustafa_m/controllers"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// postService := models.NewPostService("../learn_go/db/lenslocked_dev.db")

	// postService.AutoMigrate()

	ipAddress := getLocalIpAddress()

	port := "3000"

	fmt.Printf("Listening on %s:%s\n", ipAddress, port)

	staticC := controllers.NewStaticController()

	r.Handle("/", staticC.Homepage)

	r.Handle("/about", staticC.About)

	r.Handle("/contact", staticC.Contact)

	http.ListenAndServe(":3000", r)
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
