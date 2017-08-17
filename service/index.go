package service

import (
	"log"
	"net/http"
)

// IndexHandler handle GET /
func (s Service) IndexHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, userID := s.isAuthenticated(r)
	err := s.templates.ExecuteTemplate(w, "index", map[string]interface{}{
		"isAuthenticated": isAuthenticated,
		"userID":          userID,
	})
	if err != nil {
		log.Println(err)
	}
}
