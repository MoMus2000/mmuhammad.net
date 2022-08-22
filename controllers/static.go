package controllers

import "mustafa_m/views"

type staticController struct {
	About   *views.View
	Contact *views.View
}

func NewStaticController() *staticController {
	return &staticController{
		About:   views.NewView("bootstrap", "static/about.gohtml"),
		Contact: views.NewView("bootstrap", "static/contact.gohtml"),
	}
}
