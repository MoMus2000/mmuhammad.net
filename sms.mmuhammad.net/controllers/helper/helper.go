package helper

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
)

func ParseForm(r *http.Request, form interface{}) error {
	r.ParseForm()

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := decoder.Decode(form, r.PostForm)
	if err != nil {
		return err
	}
	return nil
}

func GetEstTime() time.Time {
	loc, _ := time.LoadLocation("EST")
	now := time.Now().In(loc)
	return now
}
