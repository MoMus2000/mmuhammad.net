package views

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

//go:embed layouts static admin home
var Res embed.FS

type View struct {
	template *template.Template
	layout   string // path to the page
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.template.ExecuteTemplate(w, v.layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := v.Render(w, nil)
	if err != nil {
		fmt.Println("Error in ServeHttp : ", err)
		panic(err)
	}
}

func NewView(layout string, files ...string) *View {
	files = append(files, getFilesFromLayout()...)
	t, err := template.ParseFS(Res, files...)
	if err != nil {
		panic(err)
	}

	return &View{
		template: t,
		layout:   layout,
	}
}

func getFilesFromLayout() []string {
	file_paths, err := getAllFilenames(&Res, ".")
	if err != nil {
		panic(err)
	}
	return file_paths
}

func getAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.Name() == "static" || entry.Name() == "admin" ||
			entry.Name() == "home" {
			continue
		}
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() && entry.Name() == "layouts" {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return
}
