package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/controllers/auth"
	"sms.mmuhammad.net/controllers/home"
	"sms.mmuhammad.net/controllers/static"
	"sms.mmuhammad.net/controllers/terminal"
	"sms.mmuhammad.net/models/db"
	"sms.mmuhammad.net/models/landing"
	"sms.mmuhammad.net/models/model_auth"
)

func main() {
	db, err := db.NewDbConnection("./db/sms_mmuhammad.db")
	if err != nil {
		panic(err)
	}

	ls := landing.NewLandingService(db)
	auth_service := model_auth.NewCreateUserService(db)

	auth_service.AutoMigrate()

	ls.AutoMigrate()

	r := mux.NewRouter()

	landC := home.NewLandingPageController(ls)

	loginC := auth.NewLoginPageController(auth_service)

	// loginC.LoginService.CreateUserService("muhammadmustafa2000@gmail.com", "Password")

	smsC := terminal.NewSmsTerminal()

	terminal.AddTerminalRoutes(r, smsC)

	home.AddHomePageRoutes(r, landC)
	auth.AddLoginRoutes(r, loginC)
	static.AddStaticRoutes(r)

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
