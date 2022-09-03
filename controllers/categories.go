package controllers

import (
	"encoding/json"
	"fmt"
	"mustafa_m/models"
	"net/http"
)

type Category struct {
	catService *models.CategoryService
}

func NewCategoryController(cs *models.CategoryService) *Category {
	return &Category{catService: cs}
}

func (cat *Category) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	cats, err := cat.catService.GetAllCategories()
	internalServerError := InternalServerError()
	if err != nil {
		fmt.Println(err)
		internalServerError.Render(w, nil)
	}
	jsonEncoding, err := json.Marshal(cats)
	fmt.Fprintln(w, string(jsonEncoding))
}