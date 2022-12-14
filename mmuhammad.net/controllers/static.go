package controllers

import (
	"mustafa_m/views"

	"github.com/gorilla/mux"
)

type staticController struct {
	About               *views.View
	PageNotFound        *views.View
	InternalServerError *views.View
	ForbiddenError      *views.View
	RandomAdvice        *views.View
}

func NewStaticController() *staticController {
	return &staticController{
		About:               views.NewView("bootstrap", "static/about.gohtml"),
		PageNotFound:        views.NewView("bootstrap", "static/404.gohtml"),
		InternalServerError: views.NewView("bootstrap", "static/500.gohtml"),
		ForbiddenError:      views.NewView("bootstrap", "static/403.gohtml"),
		RandomAdvice:        views.NewView("bootstrap", "static/randomAdvice.gohtml"),
	}
}

func AddStaticRoutes(r *mux.Router, staticC *staticController) {
	r.Handle("/about", staticC.About).Methods("GET")
	r.Handle("/randomJoke", staticC.RandomAdvice).Methods("GET")
}
