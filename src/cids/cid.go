package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	//"os"

	oidc "github.com/coreos/go-oidc"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	clientID     = "1da3e3e57166dcfd116a"                     // os.Getenv("Git_Client_Id")
	clientSecret = "806d9f4a49de69272bfaf24e5b4eb9afdebed5d9" // os.Getenv("Git_Client_Secret")
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

const appNonce = "a super secret nonce"

func ClaimNonce(nonce string) error {
	if nonce != appNonce {
		return errors.New("unrecognized nonce")
	}
	return nil
}

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx)
	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{
		ClientID:   clientID,
		ClaimNonce: ClaimNonce,
	}
	// Use the nonce source to create a custom ID Token verifier.
	nonceEnabledVerifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "https://localhost:9090/Callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	state := "foobar" // Don't do this in production.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(appNonce)), http.StatusFound)
	})

	http.HandleFunc("/Callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		// Verify the ID Token signature and nonce.
		idToken, err := nonceEnabledVerifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	log.Printf("listening on https://%s/", "localhost:9090")
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}
func contextClient(ctx context.Context) *http.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(oauth2.HTTPClient).(*http.Client); ok {
			return hc
		}
	}
	return http.DefaultClient
}

// https://accounts.google.com/o/oauth2/auth?client_id=29336734752-hkhmuft60qa4qn1ual7etmp6pfo0ib54.apps.googleusercontent.com&redirect_uri=https://account.xiaomi.com/pass/sns/verify/load&response_type=code&scope=https://www.googleapis.com/auth/userinfo.email%20https://www.googleapis.com/auth/userinfo.profile&state=7b22736964223a2270617373706f7274222c226c6f63616c65223a22656e5f494e222c226170706964223a22676f6f676c65227d
