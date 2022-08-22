package controllers

import (
	"encoding/json"
	"fmt"
	"mustafa_m/models"
	"net/http"
)

type Post struct {
	postalService *models.PostService
}

func NewPostalController(postalService *models.PostService) *Post {
	return &Post{
		postalService: postalService,
	}
}

func (post *Post) GetAllPost(w http.ResponseWriter, r *http.Request) {
	posts, err := post.postalService.GetAllPost()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	jsonEncoding, err := json.Marshal(posts)
	fmt.Fprintln(w, string(jsonEncoding))
}
