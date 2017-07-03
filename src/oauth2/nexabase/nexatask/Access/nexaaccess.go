// NexaStack user login with access token

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

var (
	config = oauth2.Config{
		ClientID:     "qwerty",
		ClientSecret: "12wsx",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:9900/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:8009/authorize",
			TokenURL: "http://localhost:8009/token",
		},
	}

	Nexastring = "chandigarh"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	u := config.AuthCodeURL(Nexastring)
	http.Redirect(w, r, u, http.StatusFound)
	fmt.Println(u)
}

func handlenexaCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.Form.Get("state")
	if state != Nexastring {
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
}

func main() {

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/oauth2", handlenexaCallback)
	fmt.Println("localhost started at 9900 ")
	log.Fatal(http.ListenAndServe(":9900", nil))

}
