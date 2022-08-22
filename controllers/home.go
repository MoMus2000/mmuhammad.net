package controllers

import "mustafa_m/views"

type Home struct {
	HomePage *views.View
}

func NewHomeController() *Home {
	return &Home{
		HomePage: views.NewView("bootstrap", "home/home.gohtml"),
	}
}
