//Github Api Oauth Personal Access
package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	// you need to generate personal access token at
	// https://github.com/settings/applications#personal-access-tokens
	personalAccessToken = "e7faa89e95e6be01eba4b038027a6e64e90ca0b4"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Printf("client.Users.Get() failed with '%s'\n", err)
		return
	}
	d, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("json.MarshlIndent() failed with %s\n", err)
		return
	}
	fmt.Printf("User:\n%s\n", string(d))
}
