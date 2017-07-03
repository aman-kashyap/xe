// id token with nonce

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
	"golang.org/x/oauth2/github"
)

var (
	clientID     = "1da3e3e57166dcfd116a"                     //  os.Getenv("Client_Id")
	clientSecret = "806d9f4a49de69272bfaf24e5b4eb9afdebed5d9" //   os.Getenv("Client_Secret")
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

	//provider, err := oidc.NewProvider(ctx, "http://localhost:6788")
	//if err != nil {
	//	log.Fatal(err)
	//}

	oidcConfig := &oidc.Config{
		ClientID:   clientID,
		ClaimNonce: ClaimNonce,
	}
	// Use the nonce source to create a custom ID Token verifier.
	nonceEnabledVerifier := github.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8888/dex/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"}, // , "gender", "birthdate"
	}

	state := "foobar" // Don't do this in production.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(appNonce)), http.StatusFound)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		url := config.AuthCodeURL(state)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
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

	log.Printf("listening on http://%s/", "localhost:9090")
	log.Fatal(http.ListenAndServe("http://localhost:9090", nil))
}

// https://api.twitter.com/oauth/authenticate?oauth_token=s3bsDQAAAAAAy-aNAAABW6ihwyc&lang=en
