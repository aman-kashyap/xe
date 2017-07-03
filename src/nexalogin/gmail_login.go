package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "nexalogin/storage"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const htmlIndex = `<html><body>
    <a href ="/login"> log in with google </a>
    </body></html>
    `

var googleClientId = "1047409585065-k9gf8lqa8ss0b0a2hug24vu2g2g7tf9m.apps.googleusercontent.com"
var googleClientSecret = "FOQSibclgP75TdDU8HNPq3c_"

var (
	googleOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:1313/GoogleCallback",
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	oauthStateString = "random"
)
var db *gorm.DB
var err error

type Mail struct {
	gorm.Model
	Email string `json:"email"`
}

type Token struct {
	gorm.Model
	Access_token string `json:"access_token"`
}

func main() {
	db := Db_connection()
	defer db.Close()
	fmt.Println("database", db)
	db.AutoMigrate(&Mail{})
	db.AutoMigrate(&Token{})

	// http.HandleFunc("/create", storage.Create)
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/GoogleCallback", handleCallback)

	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "views/home.html")
	fmt.Fprintf(w, htmlIndex)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// db := Db_connection()
	url := googleOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	db := Db_connection()
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(*token)
	fmt.Println(token.AccessToken)

	db.Create(&Token{Access_token: token.AccessToken})
	db.Save(&Token{Access_token: token.AccessToken})

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content: %s\n", contents)

	// m, err := http.Get("https://www.accounts.google.com" + m.email)
	// if err != nil {
	// 	panic(err)
	// }
	// db.Create(&Mail{Email: m.email})
	// db.Save(&Mail{Email: m.email})
}

func Db_connection() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	return db
}

// type Endpoint struct {
// 	AuthURL  string
// 	TokenURL string
// }

// var (
// 	AuthURL  = "https://accounts.google.com/o/oauth2/auth?access_type=offline"
// 	TokenURL = "https://accounts.google.com/o/oauth2/token"
// oauth2.AccessTypeOffline
// )
// AccessTypeOffline AuthCodeOption = SetAuthURLParam("access_type", "offline")
// type AuthCodeOption interface {
// 	// contains filtered or unexported methods
// }

// func SetAuthURLParam(key, value string) AuthCodeOption {
// 	return
// }
