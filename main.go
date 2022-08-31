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

	db, err := models.NewDataBaseConnection("/Users/a./Desktop/go/learn_go/db/lenslocked_dev.db")

	if err != nil {
		panic(err)
	}

	postService := models.NewPostService(db)
	adminService := models.NewAdminService(db)
	categoryService := models.NewCategoryService(db)

	postService.AutoMigrate()
	adminService.AutoMigrate()
	categoryService.AutoMigrate()

	ipAddress := getLocalIpAddress()

	port := "3000"

	fmt.Printf("Listening on %s:%s\n", ipAddress, port)

	staticC := controllers.NewStaticController()

	postalC := controllers.NewPostalController(postService)

	adminC := controllers.NewAdminController(adminService, postService, categoryService)

	homeC := controllers.NewHomeController()

	// adminService.Create(&models.Admin{Email: "muhammadmustafa4000@gmail.com",
	// 	Password: "mustafa"})

	r.Handle("/", homeC.HomePage).Methods("GET")
	r.Handle("/about", staticC.About).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.NotFoundHandler = staticC.PageNotFound
	r.MethodNotAllowedHandler = staticC.InternalServerError

	r.HandleFunc("/posts", postalC.GetAllPost).Methods("GET")
	r.HandleFunc("/admin", adminC.Login).Methods("POST")
	r.HandleFunc("/admin", adminC.GetLoginPage).Methods("GET")
	r.HandleFunc("/admin/create", adminC.GetBlogForm).Methods("GET")
	r.HandleFunc("/admin/create", adminC.SubmitBlogPost).Methods("POST")
	r.HandleFunc("/admin/delete", adminC.GetDeletePage).Methods("GET")
	r.HandleFunc("/admin/delete", adminC.SubmitDeleteRequest).Methods("POST")
	r.HandleFunc("/admin/edit", adminC.GetEditPage).Methods("GET")
	r.HandleFunc("/admin/edit", adminC.SubmitEditRequest).Methods("POST")
	r.HandleFunc("/admin/category", adminC.GetCategoryPage).Methods("GET")
	r.HandleFunc("/admin/category", adminC.SubmitCategoryFrom).Methods("POST")
	r.HandleFunc("/posts/{[a-z]+}/{[a-z]+}", postalC.GetPostFromTopic).Methods("GET")

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
