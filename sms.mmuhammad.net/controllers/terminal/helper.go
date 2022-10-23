package terminal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/xuri/excelize/v2"
)

func parseExcelForm(r *http.Request, f interface{}) (*excelize.File, error) {
	err := r.ParseMultipartForm(10000000000)
	if err != nil {
		return nil, err
	}
	file, _, err := r.FormFile("File")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	excelFile, err := excelize.OpenReader(file)

	if err != nil {
		return nil, err
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	fmt.Println(r.PostForm)
	err = decoder.Decode(f, r.PostForm)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return excelFile, nil
}
