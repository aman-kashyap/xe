//Github Api Oauth user information
package main

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
	// githuboauth
	"html/template"
	//"io/ioutil"
	//"os"
	//"encoding/json"
)

const htmlIndex = `<html><body>
	<a href ="/GithubLogin"> log in with GITHUB </a>
	</body></html>
	`

var userInfoTemplate = template.Must(template.New("").Parse(`
	<html><body>
	<p>This app is now authenticated to access your GitHub user info.</p>
	<p>User details are:</p><p>
	{{.}}
	</p>
		<p>That's it!</p>
	</body></html>
`))

var (
	oauthConf = &oauth2.Config{
		RedirectURL:  "http://localhost:8050/GithubCallback",
		ClientID:     "1a20e7f9cad317dd0cb2",
		ClientSecret: "13dbd6b367a2fd636fd494e8d2fc87a4461273e6",
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
		/*oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token"},*/
		//githuboauth.Endpoint,
	}
	oauthStateString = "random"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerMain)
	r.HandleFunc("/GithubLogin", handlerGithubLogin)
	r.HandleFunc("/GithubCallback", handlerGithubCallback)
	http.Handle("/", r)
	http.ListenAndServe(":8050", nil)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, htmlIndex)
}
func handlerGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func handlerGithubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Printf("client.Users.Get() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	buf := []string{"Github login id: ", *user.Login, "| Github email id: ", *user.Email}

	userInfoTemplate.Execute(w, buf)

	/*oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("zxdzxd")
	if err != nil {
		fmt.Printf("client.Users.Get() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}*/

	/*

		fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		response, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, "Content: %s\n", contents)*/

	/*func tokenToJSON(token *oauth2.Token) (string, error) {
	  	if d, err := json.Marshal(token); err != nil {
	  		return "", err
	  	} else {
	  		return string(d), nil
	  	}
	  }

	  func tokenFromJSON(jsonStr string) (*oauth2.Token, error) {
	  	var token oauth2.Token
	  	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
	  		return nil, err
	  	}
	  	return &token, nil
	  }*/

}
