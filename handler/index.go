package handler

import (
	"html/template"
	"net/http"
)

// Index handle GET /
func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")

	t.Execute(w, nil)
}
