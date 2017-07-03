// getting user data over cros or through html

package main

import (
	"encoding/json"
	"fmt"
	//"html/template"
	//"log"
	"net/http"
	"strings"

	"github.com/rs/cors"
)

type User struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type Sign_up struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello ashish!") // write data to response
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var sign Sign_up
	decoder.Decode(&sign)
	w.WriteHeader(http.StatusOK)

	fmt.Println("email:", sign.Email)
	fmt.Println("password:", sign.Password)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var user User
	decoder.Decode(&user)
	w.WriteHeader(http.StatusOK)

	fmt.Println("username:", user.Username)
	fmt.Println("password:", user.Password)
	fmt.Println("email:", user.Email)
	fmt.Println("passwordConfirmation:", user.PasswordConfirmation)
	fmt.Println("check")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", sayhelloName)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":9999", handler)
}

// func login(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("method:", r.Method) //get request method
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.html")
// 		t.Execute(w, nil)
// 	} else {
// 		r.ParseForm()
// 		fmt.Println("username:", r.Form["username"])
// 		fmt.Println("password:", r.Form["password"])
// 		fmt.Println("email:", r.Form["email"])
// 		fmt.Println("passwordConfirmation:", r.Form["passwordConfirmation"])
// 	}
// }
