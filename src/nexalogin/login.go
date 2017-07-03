// login page using OAUTH 2.0
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"os"
	// "strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// var (
// 	localhost = os.Getenv("localhost")
// )
var (
	oauthConf = &oauth2.Config{
		ClientID:     "1da3e3e57166dcfd116a",
		ClientSecret: "806d9f4a49de69272bfaf24e5b4eb9afdebed5d9",
		RedirectURL:  "http://172.16.0.23:9090/Callback",
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
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
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func handleCallBack(w http.ResponseWriter, r *http.Request) {
	db := Db_connection()

	r.ParseForm()
	state := r.Form.Get("state")
	if state != oauthStateString {
		fmt.Println("invalid state string. expectected %s got %s", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
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
	if err != nil && response != nil {
		fmt.Printf("Get %s \n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "http://172.16.0.12:3015/dashboard", 301)
	db.Create(&Token{Access_token: token.AccessToken})
	db.Save(&Token{Access_token: token.AccessToken})

	log.Printf("parseResponseBody: %s\n", response)
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(*token)
	//fmt.Println(json.NewEncoder(w).Decode(response))
}

func main() {

	db := Db_connection()
	defer db.Close()
	fmt.Println("database", db)
	db.AutoMigrate(&Mail{})
	db.AutoMigrate(&Token{})

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/Callback", handleCallBack)
	fmt.Println("use http://172.16.0.23:9090 to start")
	log.Fatal(http.ListenAndServe(":9090", nil))

}

func Db_connection() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	return db
}
