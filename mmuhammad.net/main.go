package main

import (
	"flag"
	"fmt"
	"log"
	"mustafa_m/controllers"
	"mustafa_m/controllers/admin"
	"mustafa_m/controllers/articles"
	"mustafa_m/controllers/fmb"
	"mustafa_m/controllers/home"
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
	fmbService := models.NewFmbService(db)

	services := []Service{
		postService,
		adminService,
		categoryService,
		monitorService,
		messageService,
		fmbService,
	}

	for _, serv := range services {
		serv.AutoMigrate()
	}

	staticC := controllers.NewStaticController()

	postalC := articles.NewPostalController(postService)

	adminC := admin.NewAdminController(adminService, postService, categoryService)

	homeC := home.NewHomeController()

	catC := articles.NewCategoryController(categoryService)

	artC := articles.NewArticlesController()

	monC := home.NewMonitorController(monitorService)

	mbC := home.NewMessageController(messageService)

	fmbC := fmb.NewTwilioController(fmbService)

	poC := home.NewPortfolioOptimization()

	r.NotFoundHandler = staticC.PageNotFound
	r.MethodNotAllowedHandler = staticC.InternalServerError

	controllers.AddStaticRoutes(r, staticC)
	controllers.AddHelperRoutes(r)
	articles.AddCategoryRoutes(r, catC)
	home.AddHomeRoutes(r, homeC)
	articles.AddPostRoutes(r, postalC)
	articles.AddArticleRoutes(r, artC)
	admin.AddAdminRoutes(r, adminC)
	home.AddMonitorRoutes(r, monC)
	home.AddMessageBoardRoutes(r, mbC)

	home.AddPORoutes(r, poC)

	fmb.AddTwilioRoutes(r, fmbC)

	// For the message board
	go home.ListenToChannel(messageService)

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
