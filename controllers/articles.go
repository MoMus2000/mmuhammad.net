package controllers

import (
	"mustafa_m/views"
	"net/http"
)

type Articles struct {
	ArticleLanding *views.View
}

func NewArticlesController() *Articles {
	return &Articles{
		ArticleLanding: views.NewView("bootstrap", "home/articleLanding.gohtml"),
	}
}

func (articles *Articles) GetArticleLandingPage(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	// Now send over the cid to the child template
	type Data struct {
		cid string
	}
	articles.ArticleLanding.Render(w, cid)
}
