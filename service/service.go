package service

import (
	"html/template"

	"github.com/gorilla/sessions"
	"github.com/tthanh/yoblog"
)

// Service represent internal service
type Service struct {
	accountStore yoblog.AccountStore
	cookieStore  *sessions.CookieStore
	templates    *template.Template
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

// SetCookieStore create new cookie store
func SetCookieStore(secret []byte) func(*Service) error {
	return func(s *Service) error {
		s.cookieStore = sessions.NewCookieStore(secret)
		return nil
	}
}
