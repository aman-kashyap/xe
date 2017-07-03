package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	//"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	//"golang.org/x/oauth2/github"
	//"golang.org/x/net/context"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "1da3e3e57166dcfd116a",
		ClientSecret: "806d9f4a49de69272bfaf24e5b4eb9afdebed5d9",
		RedirectURL:  "http://localhost:9090/Callback",
		Scopes:       []string{"user"},
		//	Endpoint:     "",
	}
	oauthStateString = "random"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

const htmlIndex = `<html><body>
	log in with <a href= "/login">Github</a>
	</body></html>
	`

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	Url, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse:", err)
	}
	parameters := url.Values{}
	parameters.Add("ClientID", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, ""))
	parameters.Add("redirect uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	Url.RawQuery = parameters.Encode()
	url := Url.String()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallBack(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Println("invalid state string. expectected %s got %s", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exachange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://api.github.com/user?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get %s \n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll %s \n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Printf("parseResponseBody: %s\n", string(response))

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/Callback", handleCallBack)
	fmt.Println("localhost started")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
