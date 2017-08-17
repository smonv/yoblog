package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/tthanh/yoblog"
	"golang.org/x/oauth2"
)

// LoginHandler handle Get /login
func (s Service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	URL, err := url.Parse(s.oauth2Config.Endpoint.AuthURL)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add("client_id", s.oauth2Config.ClientID)
	params.Add("scope", strings.Join(s.oauth2Config.Scopes, " "))
	params.Add("redirect_uri", s.oauth2Config.RedirectURL)
	params.Add("response_type", "code")
	params.Add("state", s.oauth2State)
	URL.RawQuery = params.Encode()

	http.Redirect(w, r, URL.String(), http.StatusTemporaryRedirect)
}

// CallbackHandler handle Get /callback
func (s Service) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	account, err := s.getFacebookAccount(r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	_, err = s.accountStore.GetByID(account.ID)
	if err != nil {
		_, sErr := s.accountStore.Create(account)
		if sErr != nil {
			log.Println(sErr)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	session, err := s.cookieStore.Get(r, s.cookieName)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	session.Values["user_id"] = account.ID

	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// LogoutHandler handle GET /logout
func (s Service) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	session, err := s.cookieStore.Get(r, s.cookieName)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	session.Options.MaxAge = -1

	err = s.cookieStore.Save(r, w, session)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (s Service) getFacebookAccount(r *http.Request) (*yoblog.Account, error) {
	state := r.FormValue("state")
	if state != s.oauth2State {
		return nil, errors.New("State Mismatch")
	}

	code := r.FormValue("code")

	token, err := s.oauth2Config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var account yoblog.Account
	err = json.Unmarshal(response, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s Service) isAuthenticated(r *http.Request) bool {
	session, err := s.cookieStore.Get(r, s.cookieName)
	if err != nil {
		return false
	}

	if v, ok := session.Values["user_id"]; !ok || v == "" {
		return false
	}

	return true
}
