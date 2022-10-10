package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type View struct {
	template *template.Template
	layout   string
}

func (view *View) Render(w http.ResponseWriter, data interface{}) error {
	return view.template.ExecuteTemplate(w, view.layout, data)
}

func NewView(layout string, files ...string) *View {
	for i, file := range files {
		files[i] = "./views" + file
	}
	files = append(files, getFilesFromLayout()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
	}
	return &View{
		template: t,
		layout:   layout,
	}
}

func getFilesFromLayout() []string {
	matches, _ := filepath.Glob("./views/layout/*.gohtml")
	return matches
}
