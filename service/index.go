package service

import (
	"log"
	"net/http"
)

// IndexHandler handle GET /
func (s Service) IndexHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, userID := s.isAuthenticated(r)

	posts, err := s.postStore.GetAll()
	if err != nil {
		log.Println(err)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	err = s.templates.ExecuteTemplate(w, "index", map[string]interface{}{
		"isAuthenticated": isAuthenticated,
		"userID":          userID,
		"posts":           posts,
	})
	if err != nil {
		log.Println(err)
	}
}
