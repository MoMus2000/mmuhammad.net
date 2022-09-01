package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	postalService *models.PostService
	postPage      *views.View
}

func NewPostalController(postalService *models.PostService) *Post {
	return &Post{
		postalService: postalService,
		postPage:      views.NewView("article", "static/content/post.gohtml"),
	}
}

func (post *Post) GetAllPost(w http.ResponseWriter, r *http.Request) {
	posts, err := post.postalService.GetAllPost()
	internalServerError := InternalServerError()
	if err != nil {
		fmt.Println(err)
		internalServerError.Render(w, nil)
	}
	jsonEncoding, err := json.Marshal(posts)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (post *Post) GetPostFromTopic(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// Provide the render page with the data from the database to create the
	// article and view it
	internalServerError := InternalServerError()
	arr := strings.Split(r.URL.Path, "/")
	id := arr[len(arr)-2]
	idToUint, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		internalServerError.Render(w, nil)
	}

	type Data struct {
		Topic   string
		Summary string
		Content template.HTML
	}

	data := post.postalService.GetPost(uint(idToUint))

	dataObject := &Data{
		Topic:   data.Topic,
		Summary: data.Summary,
		Content: template.HTML(data.Content)}

	err = post.postPage.Render(w, dataObject)
	if err != nil {
		internalServerError.Render(w, nil)
	}
}

func (post *Post) GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am finally here")
	cid := r.URL.Query().Get("cid")
	posts, err := post.postalService.GetAllPostByCategory(cid)
	internalServerError := InternalServerError()
	if err != nil {
		fmt.Println(err)
		internalServerError.Render(w, nil)
	}

	jsonEncoding, err := json.Marshal(posts)
	fmt.Fprintln(w, string(jsonEncoding))
}
