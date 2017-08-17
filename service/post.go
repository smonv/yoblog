package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tthanh/yoblog"
)

func (s Service) AccountPostsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	accountID := vars["aid"]

	posts, err := s.postStore.GetByOwnerID(accountID)
	if err != nil {
		log.Println(err)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	isAuthenticated, userID := s.isAuthenticated(r)

	err = s.templates.ExecuteTemplate(w, "posts", map[string]interface{}{
		"isAuthenticated": isAuthenticated,
		"userID":          userID,
		"posts":           posts,
	})
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func (s Service) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, userID := s.isAuthenticated(r)

	err := s.templates.ExecuteTemplate(w, "new_post", map[string]interface{}{
		"isAuthenticated": isAuthenticated,
		"userID":          userID,
	})
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func (s Service) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	_, userID := s.isAuthenticated(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	_userID := r.FormValue("user_id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	post := &yoblog.Post{
		OwnerID: _userID,
		Title:   title,
		Content: content,
	}

	_, err = s.postStore.Create(post)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/accounts/"+userID+"/posts", http.StatusSeeOther)
}
