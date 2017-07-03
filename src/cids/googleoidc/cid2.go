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
)

var (
	clientID     = "1047409585065-k9gf8lqa8ss0b0a2hug24vu2g2g7tf9m.apps.googleusercontent.com" //  os.Getenv("Client_Id")
	clientSecret = "FOQSibclgP75TdDU8HNPq3c_"                                                  //   os.Getenv("Client_Secret")
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

	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
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
		RedirectURL:  "http://localhost:1313/GoogleCallback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"}, // , "gender", "birthdate"
	}

	state := "foobar" // Don't do this in production.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(appNonce)), http.StatusFound)
	})

	http.HandleFunc("/GoogleCallback", func(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("listening on http://%s/", "localhost:1313")
	log.Fatal(http.ListenAndServe("localhost:1313", nil))
}

// https://api.twitter.com/oauth/authenticate?oauth_token=s3bsDQAAAAAAy-aNAAABW6ihwyc&lang=en
