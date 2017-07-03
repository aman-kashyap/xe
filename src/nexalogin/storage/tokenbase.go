// getting user data from html/http and saving to database

package storage

import (
	"encoding/json"
	"fmt"
	// "html/template"
	// "database/sql"
	// "log"
	"net/http"
	//"strings"

	// "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func Create(w http.ResponseWriter, r *http.Request) {
	db := Db_connection()
	if r.Method != "POST" {
		return
	}
	if r.URL.Path != "/GoogleCallback" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var m Mail
	var t Token

	decoder.Decode(&m)
	decoder.Decode(&t)

	// db.Create(&Product{Code: d.Code, Price: d.Price})
	db.Create(&Mail{Email: m.Email})
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", m.Email)
	// mail := Mail{email: "aman@xenondigilabs.com"}

	db.Create(&Token{Access_token: t.Access_token})
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", t.Access_token)
}
func Update(w http.ResponseWriter, r *http.Request) {
	db := Db_connection()
	if r.Method != "POST" {
		return
	}
	if r.URL.Path != "/GoogleCallback" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var m Mail
	var t Token
	decoder.Decode(&m)
	decoder.Decode(&t)

	// db.Table("mails").Where("id = ?", 1).Update("price", d.Price)
	db.Save(&Mail{Email: m.Email})
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", m.Email)
	db.Save(&Token{Access_token: t.Access_token})
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", t.Access_token)
}

func Db_connection() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	return db
}

// func init() {
// 	db := Db_connection()
// 	defer db.Close()
// 	db.AutoMigrate(&Mail{})
// 	db.AutoMigrate(&Token{})
// }

// func main() {
// 	// db := Db_connection()
// 	// defer db.Close()
// 	// db.AutoMigrate(&Mail{})
// 	// db.AutoMigrate(&Token{})
// 	// db.Create(&Mail{})
// 	// db.Create(&Token{})

// 	r := mux.NewRouter()
// 	r.HandleFunc("/create", Create)
// 	http.Handle("/", r)
// 	log.Fatal(http.ListenAndServe(":9009", nil))
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	db := Db_connection()
// 	if r.Method == "GET" {
// 		// t, _ := template.ParseFiles("login.html")
// 		// t.Execute(w, nil)

// 		var email = Mail{}
// 		r.Form.Get("email")
// 		err := db.Create(&email)
// 		if err != nil {
// 			panic(err)
// 		}
// 	} else {
// 		r.ParseForm()
// 		// var email = Email{}
// 		// r.Form.Get("email")
// 		// err := db.Create(&email)
// 		// if err != nil {
// 		// 	panic(err)
// 		// }
// 		// logic part of log in
// 		// fmt.Println("username:", r.Form["username"])
// 		// fmt.Println("email:", r.Form["email"])
// 	}
// }
