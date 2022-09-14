package main

import (
	"fmt"
	"log"
	"mustafa_m/controllers"
	"mustafa_m/models"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	db, err := models.NewDataBaseConnection("./db/lenslocked_dev.db")

	if err != nil {
		panic(err)
	}

	postService := models.NewPostService(db)
	adminService := models.NewAdminService(db)
	categoryService := models.NewCategoryService(db)
	monitorService := models.NewMonitorService(db)

	postService.AutoMigrate()
	adminService.AutoMigrate()
	monitorService.AutoMigrate()
	categoryService.AutoMigrate()

	ipAddress := getLocalIpAddress()

	port := "3000"

	fmt.Printf("Listening on %s:%s\n", ipAddress, port)

	staticC := controllers.NewStaticController()

	postalC := controllers.NewPostalController(postService)

	adminC := controllers.NewAdminController(adminService, postService, categoryService)

	homeC := controllers.NewHomeController()

	catC := controllers.NewCategoryController(categoryService)

	artC := controllers.NewArticlesController()

	monC := controllers.NewMonitorController(monitorService)

	r.Handle("/about", staticC.About).Methods("GET")

	r.NotFoundHandler = staticC.PageNotFound
	r.MethodNotAllowedHandler = staticC.InternalServerError

	r.HandleFunc("/", controllers.WrapIPHandler(homeC.GetHomePage)).Methods("GET")
	r.HandleFunc("/posts", postalC.GetAllPost).Methods("GET")
	r.HandleFunc("/categories", catC.GetAllCategories).Methods("GET")
	r.HandleFunc("/admin", adminC.Login).Methods("POST")
	r.HandleFunc("/admin", adminC.GetLoginPage).Methods("GET")
	r.HandleFunc("/admin/create", adminC.GetBlogForm).Methods("GET")
	r.HandleFunc("/admin/create", adminC.SubmitBlogPost).Methods("POST")
	r.HandleFunc("/admin/delete", adminC.GetDeletePage).Methods("GET")
	r.HandleFunc("/admin/delete", adminC.SubmitArticleDeleteRequest).Methods("POST")
	r.HandleFunc("/admin/edit", adminC.GetEditPage).Methods("GET")
	r.HandleFunc("/admin/edit", adminC.SubmitArticleEditRequest).Methods("POST")
	r.HandleFunc("/admin/category", adminC.GetCategoryPage).Methods("GET")
	r.HandleFunc("/admin/category", adminC.SubmitCategoryFrom).Methods("POST")
	r.HandleFunc("/admin/category/edit", adminC.GetCategoryEditPage).Methods("GET")
	r.HandleFunc("/admin/category/edit", adminC.SubmitCategoryEditRequest).Methods("POST")
	r.HandleFunc("/admin/category/delete", adminC.GetCategoryDeletePage).Methods("GET")
	r.HandleFunc("/admin/category/delete", adminC.SubmitCategoryDeleteRequest).Methods("POST")
	r.HandleFunc("/posts/{[a-z]+}/{[a-z]+}", postalC.GetPostFromTopic).Methods("GET")

	r.HandleFunc("/articles", artC.GetArticleLandingPage).Methods("GET")
	r.HandleFunc("/postByCat", postalC.GetPostsByCategory).Methods("GET")
	r.HandleFunc("/signout", adminC.SignoutJWT).Methods("GET")

	r.Handle("/market", monC.MonitorPage).Methods("GET")
	r.HandleFunc("/usopen", monC.GetUsdToPkr).Methods("GET")
	r.HandleFunc("/steel", monC.GetSteelRates).Methods("GET")
	r.HandleFunc("/oil", monC.GetOilRates).Methods("GET")

	r.HandleFunc("/api/v1/monitoring/basement", monC.GetBasementRates).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/apartment", monC.GetApartmentRates).Methods("GET")

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
