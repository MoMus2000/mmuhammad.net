package admin

import (
	"fmt"
	"mustafa_m/controllers"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Admin struct {
	LoginPage          *views.View
	AdminService       *models.AdminService
	PostService        *models.PostService
	CategoryService    *models.CategoryService
	BlogForm           *views.View
	DeleteForm         *views.View
	EditForm           *views.View
	CategoryForm       *views.View
	CategoryDeleteForm *views.View
	CategoryEditForm   *views.View
}

func NewAdminController(adminService *models.AdminService, ps *models.PostService,
	cs *models.CategoryService) *Admin {
	return &Admin{
		LoginPage:          views.NewView("bootstrap", "admin/login.gohtml"),
		AdminService:       adminService,
		PostService:        ps,
		CategoryService:    cs,
		BlogForm:           views.NewView("bootstrap", "admin/post/blogForm.gohtml"),
		DeleteForm:         views.NewView("bootstrap", "admin/post/deleteForm.gohtml"),
		EditForm:           views.NewView("bootstrap", "admin/post/editForm.gohtml"),
		CategoryForm:       views.NewView("bootstrap", "admin/category/categoryForm.gohtml"),
		CategoryDeleteForm: views.NewView("bootstrap", "admin/category/categoryDeleteForm.gohtml"),
		CategoryEditForm:   views.NewView("bootstrap", "admin/category/categoryEditForm.gohtml"),
	}
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type BlogForm struct {
	Topic      string `schema:"Topic"`
	Summary    string `schema:"Summary"`
	Imgur_URL  string `schema:"Imgur"`
	Content    string
	CategoryId string `schema:"CID"`
}

type DeleteForm struct {
	Id string `schema:"Id"`
}

type EditForm struct {
	Id        string `schema:"ID"`
	Topic     string `schema:"Topic"`
	Summary   string `schema:"Summary"`
	Imgur_URL string `schema:"Imgur"`
}

type CategoryForm struct {
	Category  string `schema:"Cat"`
	Summary   string `schema:"Summary"`
	Imgur_URL string `schema:"Imgur"`
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (admin *Admin) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request recieved")
	internalServerError := controllers.InternalServerError()
	form := LoginForm{}
	controllers.ParseForm(r, &form)
	adminTemp := models.Admin{Email: form.Email, Password: form.Password}
	result, err := admin.AdminService.ByEmail(&adminTemp)
	if err != nil {
		fmt.Println(err)
		internalServerError.Render(w, nil)
	}
	fmt.Println(result)

	createJWT(w, &adminTemp)

	http.Redirect(w, r, "/admin/create", http.StatusFound)
}

func (admin *Admin) SubmitBlogPost(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	internalServerError := controllers.InternalServerError()
	form := BlogForm{}
	content, err := controllers.ParseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	form.Content = content
	// TODO: Change to take in the IMGUR URL and the uploaded file
	// Save the file content into the database
	post := models.Post{Topic: form.Topic, Content: form.Content, Summary: form.Summary,
		Imgur_URL: form.Imgur_URL, Date: time.Now().String(), CategoryId: form.CategoryId}
	err = admin.PostService.CreatePost(&post)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	http.Redirect(w, r, "/admin/create", http.StatusFound)
}

func (admin *Admin) SubmitArticleDeleteRequest(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	form := DeleteForm{}
	internalServerError := controllers.InternalServerError()
	_, err := controllers.ParseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	fmt.Println(form)
	idToUint, err := strconv.ParseUint(form.Id, 0, 64)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	err = admin.PostService.DeletePost(uint(idToUint))
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	http.Redirect(w, r, "/admin/delete", http.StatusFound)
}

func (admin *Admin) SubmitArticleEditRequest(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	form := EditForm{}
	internalServerError := controllers.InternalServerError()
	content, err := controllers.ParseForm(r, &form)
	if err != nil {
		fmt.Println("Error is here 2", err)
		internalServerError.Render(w, nil)
		return
	}
	post := models.Post{Topic: form.Topic, Summary: form.Summary,
		Imgur_URL: form.Imgur_URL, Content: content}
	err = admin.AdminService.UpdateChangesFromEdit(&post, form.Id)
	if err != nil {
		fmt.Println("Error I got back ", err)
		internalServerError.Render(w, nil)
		return
	}
	http.Redirect(w, r, "/admin/edit", http.StatusFound)
}

func (admin *Admin) SubmitCategoryDeleteRequest(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	form := DeleteForm{}
	internalServerError := controllers.InternalServerError()
	_, err := controllers.ParseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	fmt.Println(form)
	idToUint, err := strconv.ParseUint(form.Id, 0, 64)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	err = admin.CategoryService.DeleteCategory(idToUint)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	http.Redirect(w, r, "/admin/category/delete", http.StatusFound)
}

func (admin *Admin) SubmitCategoryFrom(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	internalServerError := controllers.InternalServerError()
	form := CategoryForm{}
	_, err := controllers.ParseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	cat := models.Category{CategoryName: form.Category, Imgur_URL: form.Imgur_URL,
		CategorySummary: form.Summary, CreationDate: time.Now().String()}
	err = admin.CategoryService.Create(&cat)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	http.Redirect(w, r, "/admin/category", http.StatusFound)
}

func (admin *Admin) SubmitCategoryEditRequest(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	form := EditForm{}
	internalServerError := controllers.InternalServerError()
	_, err := controllers.ParseForm(r, &form)
	if err != nil {
		fmt.Println("Error is here 2", err)
		internalServerError.Render(w, nil)
		return
	}

	fmt.Println("GOGOGOGOGOG")

	category := models.Category{CategoryName: form.Topic, CategorySummary: form.Summary,
		Imgur_URL: form.Imgur_URL}

	err = admin.CategoryService.UpdateChangesCategoryFromEdit(&category, form.Id)
	if err != nil {
		fmt.Println("Error I got back ", err)
		internalServerError.Render(w, nil)
		return
	}
	http.Redirect(w, r, "/admin/category/edit", http.StatusFound)
}

func (admin *Admin) GetBlogForm(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{LoggedIn: "true"}
	admin.BlogForm.Render(w, data)
}

func (admin *Admin) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	if controllers.ValidateJWT(r) {
		http.Redirect(w, r, "/admin/create", http.StatusFound)
		return
	}
	admin.LoginPage.Render(w, nil)
}

func (admin *Admin) GetDeletePage(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{LoggedIn: "true"}
	admin.DeleteForm.Render(w, data)
}

func (admin *Admin) GetEditPage(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{LoggedIn: "true"}
	admin.EditForm.Render(w, data)
}

func (admin *Admin) GetCategoryPage(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{LoggedIn: "true"}
	admin.CategoryForm.Render(w, data)
}

func (admin *Admin) GetCategoryDeletePage(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{LoggedIn: "true"}
	admin.CategoryDeleteForm.Render(w, data)
}

func (admin *Admin) GetCategoryEditPage(w http.ResponseWriter, r *http.Request) {
	if !controllers.ValidateJWT(r) {
		controllers.ForbiddenError().Render(w, nil)
		return
	}
	data := &views.Data{LoggedIn: "true"}
	admin.CategoryEditForm.Render(w, data)
}

func AddAdminRoutes(r *mux.Router, adminC *Admin) {
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
	r.HandleFunc("/signout", adminC.SignoutJWT).Methods("GET")
}
