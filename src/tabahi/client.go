package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

const appNonce = "a super secret nonce"

func ClaimNonce(nonce string) error {
	if nonce != appNonce {
		return errors.New("unrecognized nonce")
	}
	return nil
}

func main() {

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:9080/log")
	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{
		ClientID:   "qwerty",
		ClaimNonce: ClaimNonce,
	}

	nonceEnabledVerifier := provider.Verifier(oidcConfig)

	config = oauth2.Config{
		ClientID:     "qwerty",
		ClientSecret: "12wsx",
		Scopes:       []string{oidc.ScopeOpenID, "all"},
		RedirectURL:  "http://localhost:9080/oauth2",
		provider.Endpoint(),
	}

	state := "foobar"

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/oauth2", handleAuth)
	http.HandleFunc("/log", handleLogin)

	log.Println("Client is running at 9080 port.")
	log.Fatal(http.ListenAndServe(":9080", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	// u := config.AuthCodeURL("xyz")
	http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(appNonce)), http.StatusFound)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/home.html")
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
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
}
