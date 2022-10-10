package articles

import (
	"encoding/json"
	"fmt"
	"mustafa_m/controllers"
	"mustafa_m/models"
	"net/http"

	"github.com/gorilla/mux"
)

type Category struct {
	catService *models.CategoryService
}

func NewCategoryController(cs *models.CategoryService) *Category {
	return &Category{catService: cs}
}

func (cat *Category) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	cats, err := cat.catService.GetAllCategories()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		fmt.Println(err)
		internalServerError.Render(w, nil)
	}
	jsonEncoding, err := json.Marshal(cats)
	fmt.Fprintln(w, string(jsonEncoding))
}

func AddCategoryRoutes(r *mux.Router, catC *Category) {
	r.HandleFunc("/api/v1/categories", catC.GetAllCategories).Methods("GET")
}
