package controllers

import "mustafa_m/views"

type staticController struct {
	About               *views.View
	PageNotFound        *views.View
	InternalServerError *views.View
	ForbiddenError      *views.View
}

func NewStaticController() *staticController {
	return &staticController{
		About:               views.NewView("bootstrap", "static/about.gohtml"),
		PageNotFound:        views.NewView("bootstrap", "static/404.gohtml"),
		InternalServerError: views.NewView("bootstrap", "static/500.gohtml"),
		ForbiddenError:      views.NewView("bootstrap", "static/403.gohtml"),
	}
}
