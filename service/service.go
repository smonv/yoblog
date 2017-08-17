package service

import (
	"html/template"

	"github.com/gorilla/sessions"
	"github.com/tthanh/yoblog"
	"golang.org/x/oauth2"
)

// Service represent internal service
type Service struct {
	accountStore yoblog.AccountStore
	postStore    yoblog.PostStore

	cookieStore *sessions.CookieStore
	cookieName  string

	oauth2Config *oauth2.Config
	oauth2State  string

	templates *template.Template
}

// New create new service
func New(options ...func(*Service) error) (*Service, error) {
	var err error

	service := Service{}

	for _, option := range options {
		err = option(&service)

		if err != nil {
			return nil, err
		}
	}

	service.templates, err = template.ParseFiles(
		"view/header.tmpl",
		"view/footer.tmpl",
		"view/index.html",
		"view/posts.html",
		"view/new_post.html",
	)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

// SetAccountStore create new account store
func SetAccountStore(accountStore yoblog.AccountStore) func(*Service) error {
	return func(s *Service) error {
		s.accountStore = accountStore
		return nil
	}
}

func SetPostStore(postStore yoblog.PostStore) func(s *Service) error {
	return func(s *Service) error {
		s.postStore = postStore
		return nil
	}
}

// SetCookieStore create new cookie store
func SetCookieStore(secret []byte) func(*Service) error {
	return func(s *Service) error {
		s.cookieStore = sessions.NewCookieStore(secret)
		return nil
	}
}

// SetCookieName set name for cookie
func SetCookieName(name string) func(*Service) error {
	return func(s *Service) error {
		s.cookieName = name
		return nil
	}
}

// SetOAuth2Config ...
func SetOAuth2Config(cfg *oauth2.Config) func(*Service) error {
	return func(s *Service) error {
		s.oauth2Config = cfg
		return nil
	}
}

// SetOAuth2State ...
func SetOAuth2State(state string) func(*Service) error {
	return func(s *Service) error {
		s.oauth2State = state
		return nil
	}
}
