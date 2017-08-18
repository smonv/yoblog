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

		return
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

func (s Service) ViewPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID := vars["pid"]

	isAuthenticated, userID := s.isAuthenticated(r)

	post, err := s.postStore.GetByID(postID)
	if err != nil {
		log.Println(err)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	comments, err := s.postStore.GetPostComments(post.ID)
	if err != nil {
		log.Println(err)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	err = s.templates.ExecuteTemplate(w, "post", map[string]interface{}{
		"isAuthenticated": isAuthenticated,
		"userID":          userID,
		"post":            post,
		"comments":        comments,
	})
	if err != nil {
		log.Println(err)
	}
}

func (s Service) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID := vars["pid"]

	err := r.ParseForm()
	if err != nil {
		log.Println(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	_userID := r.FormValue("user_id")
	content := r.FormValue("content")

	comment := &yoblog.Comment{
		OwnerID: _userID,
		PostID:  postID,
		Content: content,
	}

	_, err = s.postStore.CreateComment(comment)
	if err != nil {
		log.Panicln(err)
	}

	http.Redirect(w, r, "/posts/"+postID, http.StatusSeeOther)
}
