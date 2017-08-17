package service

import (
	"html/template"
	"net/http"
)

// IndexHandler handle GET /
func (s Service) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")

	t.Execute(w, nil)
}
