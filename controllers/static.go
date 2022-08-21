package controllers

import "mustafa_m/views"

type staticController struct {
	Homepage *views.View
	About    *views.View
	Contact  *views.View
}

func NewStaticController() *staticController {
	return &staticController{
		Homepage: views.NewView("bootstrap", "static/home.gohtml"),
		About:    views.NewView("bootstrap", "static/about.gohtml"),
		Contact:  views.NewView("bootstrap", "static/contact.gohtml"),
	}
}
