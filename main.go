package main

import (
	"flag"
	"fmt"
	"log"
	"mustafa_m/controllers"
	"mustafa_m/models"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	AutoMigrate() error
}

func main() {
	controllers.IsProduction = controllers.CheckProduction()
	flag.Parse()

	db, err := models.NewDataBaseConnection("./db/lenslocked_dev.db")

	if *controllers.IsProduction {
		fmt.Println("In production mode ..")
	} else {
		db.LogMode(true)
		fmt.Println("In development mode ..")
	}

	r := mux.NewRouter().StrictSlash(true)

	if err != nil {
		panic(err)
	}

	postService := models.NewPostService(db)
	adminService := models.NewAdminService(db)

	categoryService := models.NewCategoryService(db)
	monitorService := models.NewMonitorService(db)
	messageService := models.NewMessageService(db)

	services := []Service{
		postService,
		adminService,
		categoryService,
		monitorService,
		messageService,
	}

	for _, serv := range services {
		serv.AutoMigrate()
	}

	staticC := controllers.NewStaticController()

	postalC := controllers.NewPostalController(postService)

	adminC := controllers.NewAdminController(adminService, postService, categoryService)

	homeC := controllers.NewHomeController()

	catC := controllers.NewCategoryController(categoryService)

	artC := controllers.NewArticlesController()

	monC := controllers.NewMonitorController(monitorService)

	mbC := controllers.NewMessageController(messageService)

	fmbC := controllers.NewTwilioController(adminService)

	r.NotFoundHandler = staticC.PageNotFound
	r.MethodNotAllowedHandler = staticC.InternalServerError

	controllers.AddStaticRoutes(r, staticC)
	controllers.AddHelperRoutes(r)
	controllers.AddCategoryRoutes(r, catC)
	controllers.AddHomeRoutes(r, homeC)
	controllers.AddPostRoutes(r, postalC)
	controllers.AddArticleRoutes(r, artC)
	controllers.AddAdminRoutes(r, adminC)
	controllers.AddMonitorRoutes(r, monC)
	controllers.AddMessageBoardRoutes(r, mbC)

	controllers.AddTwilioRoutes(r, fmbC)

	// For the message board
	go controllers.ListenToChannel(messageService)

	ipAddress := getLocalIpAddress()

	port := "3000"

	fmt.Printf("Listening on %s:%s\n", ipAddress, port)

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
