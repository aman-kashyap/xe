//
package main

import (
	"fmt"
	//"html"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
)

const htmlIndex = `<html><body>
	<a href ="/GoogleLogin"> log in with google </a>
	</body></html>
	`

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerMain)
	r.HandleFunc("/GoogleLogin", handlerGoogleLogin)
	r.HandleFunc("/GoogleCallback", handlerGoogleCallback)
	http.Handle("/", r)
	http.ListenAndServe(":1313", nil)
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:1313/GoogleCallback",

		ClientID: "1047409585065-k9gf8lqa8ss0b0a2hug24vu2g2g7tf9m.apps.googleusercontent.com",

		ClientSecret: "FOQSibclgP75TdDU8HNPq3c_",
		Scopes: []string{"https://www.googleapis.com/auth/plus.login",
			"https://www.googleapis.com/auth/plus.me",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	oauthStateString = "random"
)

func handlerMain(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, htmlIndex)
}
func handlerGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func handlerGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content: %s\n", contents)

}
