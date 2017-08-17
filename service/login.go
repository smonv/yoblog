package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/tthanh/yoblog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "1943526969221067",
		ClientSecret: "226c84c4f865b0844956c4359e611d03",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
	oauthStateString = "thisshouldberandom"
)

// LoginHandler handle Get /login
func (s Service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add("client_id", oauthConf.ClientID)
	params.Add("scope", strings.Join(oauthConf.Scopes, " "))
	params.Add("redirect_uri", oauthConf.RedirectURL)
	params.Add("response_type", "code")
	params.Add("state", oauthStateString)
	URL.RawQuery = params.Encode()

	http.Redirect(w, r, URL.String(), http.StatusTemporaryRedirect)
}

// CallbackHandler handle Get /callback
func (s Service) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")

	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var account yoblog.Account
	err = json.Unmarshal(response, &account)
	if err != nil {
		log.Panicln(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	_, err = s.accountStore.GetByID(account.ID)
	if err != nil {
		_, sErr := s.accountStore.Create(&account)
		if sErr != nil {
			log.Println(sErr)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	session, err := s.cookieStore.Get(r, "yoblog")
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

	session, err := s.cookieStore.Get(r, "yoblog")
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

func (s Service) isAuthenticated(r *http.Request) bool {
	session, err := s.cookieStore.Get(r, "yoblog")
	if err != nil {
		return false
	}

	if v, ok := session.Values["user_id"]; !ok || v == "" {
		return false
	}

	return true
}
