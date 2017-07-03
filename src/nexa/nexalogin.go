package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	config = oauth2.Config{
		ClientID:     "qwerty",
		ClientSecret: "12wsx",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:9090",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:8000/authorize",
			TokenURL: "http://localhost:8000/token",
		},
	}
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handleMain)
	r.HandleFunc("/login", handleLogin)
	r.HandleFunc("/Callback", handleCallback)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := config.AuthCodeURL("xyz")
		http.Redirect(w, r, u, http.StatusFound)
	})

	r.HandleFunc("/oauth2", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		state := r.Form.Get("state")
		if state != "xyz" {
			http.Error(w, "State invalid", http.StatusBadRequest)
			return
		}
		code := r.Form.Get("code")
		if code == "" {
			http.Error(w, "Code not found", http.StatusBadRequest)
			return
		}
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(*token)
	})

	r.HandleFunc("/redirect", handleAuth)
	http.Handle("/", r)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "localhost:8000", http.StatusTemporaryRedirect)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := githubOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := githubOAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(token.AccessToken)
	response, err := http.Get("https://api.github.com/user?access_token=" + url.QueryEscape(token.AccessToken))
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content: %s\n", contents)
}
