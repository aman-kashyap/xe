// twitter OAuth2 login
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

// var (
// 	requestURL      = "https://api.twitter.com/oauth/request_token" // Request token URL
// 	authorizeURL    = "https://api.twitter.com/oauth/authorize"     // Authorize URL
// 	authenticateURL = "https://api.twitter.com/oauth/authenticate"
// 	tokenURL        = "https://api.twitter.com/oauth/access_token" // Access token URL
// 	endpointProfile = "https://api.twitter.com/1.1/account/verify_credentials.json"
// )

var (
	oauthConf = &oauth2.Config{
		ConsumerKey:    "5Qf16mzNd0IQ5ggAQSuOdtMvG",
		ConsumerSecret: "iXa3Of8k21FxTQp5KyBGZgdqc5APvAlKIgTp1QzN27v5c5qBmt",
		RedirectURL:    "http://127.0.0.1:8998/Callback",
		// Scopes:         []string{"read-write"},
		// Endpoint:       oauth2.Endpoint,
		// {
		// 	AuthURL:  "https://api.twitter.com/oauth/authorize",
		// 	TokenURL: "https://api.twitter.com/oauth/request_token",
		// },
	}
	// oauthStateString = "random"
)

const htmlIndex = `<html><body>
	log in with <a href= "/login">Twitter</a>
	</body></html>
	`

func handleMain(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(htmlIndex))
	fmt.Fprintf(w, htmlIndex)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func handleCallBack(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.Form.Get("state")
	if state != oauthStateString {
		fmt.Println("invalid state string. expectected %s got %s", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://api.twitter.com/oauth2/token" + url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get %s \n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil && response != nil {
		fmt.Printf("Get %s \n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	log.Printf("parseResponseBody: %s\n", response)
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(*token)
	//fmt.Println(json.NewEncoder(w).Decode(response))
}
func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/Callback", handleCallBack)
	fmt.Println("localhost started")
	log.Fatal(http.ListenAndServe(":8998", nil))

}
