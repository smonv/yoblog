package service

import (
	"net/http"
)

// IndexHandler handle GET /
func (s Service) IndexHandler(w http.ResponseWriter, r *http.Request) {
	s.templates.ExecuteTemplate(w, "index", ViewData{
		IsAuthenticated: s.isAuthenticated(r),
	})
}
