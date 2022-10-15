package static

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddStaticRoutes(r *mux.Router) {
	// Linking static css, img and js
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("views/layout/style/"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/",
		http.FileServer(http.Dir("views/layout/style/img/"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
		http.FileServer(http.Dir("views/js/"))))
}
