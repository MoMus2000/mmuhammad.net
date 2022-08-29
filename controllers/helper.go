package controllers

import "mustafa_m/views"

func InternalServerError() *views.View {
	return NewStaticController().InternalServerError
}
