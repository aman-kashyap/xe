package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

var (
	config = oauth2.Config{
		ClientID:     "1da3e3e57166dcfd116a",
		ClientSecret: "806d9f4a49de69272bfaf24e5b4eb9afdebed5d9",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:9090/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:9090/authorize",
			TokenURL: "http://localhost:9090/token",
		},
	}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := config.AuthCodeURL("xyz")
		http.Redirect(w, r, u, http.StatusFound)
	})

	http.HandleFunc("/oauth2", func(w http.ResponseWriter, r *http.Request) {
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

	log.Println("Client is running at 9090 port.")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
