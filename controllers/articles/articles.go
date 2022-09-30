package articles

import (
	"fmt"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/mux"
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
	offset := r.URL.Query().Get("offset")
	fmt.Println(offset)
	// Now send over the cid to the child template
	type Data struct {
		LoggedIn    string
		cid         string
		offset      string
		FmbLoggedIn string
	}
	data := &Data{cid: cid}
	err := articles.ArticleLanding.Render(w, data)
	fmt.Println(err)
}

func AddArticleRoutes(r *mux.Router, artC *Articles) {
	r.HandleFunc("/articles", artC.GetArticleLandingPage).Methods("GET")
}
