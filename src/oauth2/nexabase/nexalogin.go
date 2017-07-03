// getting user data from html/http and saving to database

package main

import (
	"encoding/json"
	// "fmt"
	// "html/template"
	// "database/sql"
	"log"
	"net/http"
	//"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// var db *gorm.DB
var err error

type User struct {
	Emails   []Email   `json:"emails"`
	Accounts []Account `json:"accounts"`
}

type Email struct {
	ID    int    `json:"id"`
	Email string `json:"email",gorm:"type:varchar(100);unique_index"`
}

type Account struct {
	Github   bool `json:"github"`
	Gmail    bool `json:"gmail"`
	Twitter  bool `json:"twitter"`
	Nexamail bool `json:"nexamail"`
}

type Data struct {
	Input User `json:"input"`
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("email:", r.Form["Email"])
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	db := Db_connection()
	if r.Method != "POST" {
		return
	}

	if r.URL.Path != "/login" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var d Product

	decoder.Decode(&d)
	db.Create(&Product{Code: d.Code, Price: d.Price})
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", d.Price, d.Code)
	email := r.FormValue("Email")
	// decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)
	e := &Email{email}
	// a := Account{}
	// usr := User{Emails: []Email{}, Accounts: []Account{}}
	// x := Data{Input: usr}
	http.ServeFile(w, r, "login.html")
	decoder.Decode(&e)
	db.Create(&e)
	// db.Create(&x)
}

func Db_connection() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	return db
}

func main() {
	db := Db_connection()
	defer db.Close()

	// db.AutoMigrate(&User{})
	db.AutoMigrate(&Email{})
	// db.AutoMigrate(&Account{})
	// db.AutoMigrate(&Data{})

	// db.Create(&Email{Email: "aman@xenondigilabs.com"})
	// db.Create(&User{Emails: []Email{}, Accounts: []Account{}})
	// db.Create(&Account{})
	// db.Create(&Data{})

	// Read
	// var email Email
	// db.First(&email, 1)                                     // find product with id 1
	// db.First(&email, "email = $", "aman@xenondigilabs.com") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)

	r := mux.NewRouter()
	r.HandleFunc("/login", CreateHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9009", nil))
}

// db := Db_connection()

// func check(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func CreateHandler(w http.ResponseWriter, r *http.Request) {
// 	db := Db_connection()
// 	decoder := json.NewDecoder(r.Body)
// }

// func Db_connection() (db *gorm.DB) {
// 	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 sslmode=disable password=postgres")
// 	check(err)
// 	db.LogMode(false)
// 	return db
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	fmt.Println("method:", r.Method) //get request method
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.html")
// 		t.Execute(w, nil)
// 	} else {
// 		r.ParseForm()
// 		fmt.Println("email:", r.Form["Email"])
// 	}
// }

// user := Email{Email: "aman@xenondigilabs.com"}
// db.NewRecord(user)
// db.Create(&user)

// r.HandleFunc("/delete", DeleteHandler)
// r.HandleFunc("/update", UpdateHandler)
// r.HandleFunc("/read", ReadHandler)
// r.HandleFunc("/", sayhelloName)

// func DeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	var product Product
// 	db := Db_connection()
// 	fmt.Println("delete")
// 	if r.Method != "POST" {
// 		return
// 	}
// 	if r.URL.Path != "/delete" {
// 		return
// 	}

// 	// Delete - delete product
// 	db.Delete(&product)
// }
// func UpdateHandler(w http.ResponseWriter, r *http.Request) {
// 	//var product Product
// 	db := Db_connection()
// 	fmt.Println("update")
// 	if r.Method != "POST" {
// 		return
// 	}
// 	if r.URL.Path != "/update" {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	var d Product

// 	decoder.Decode(&d)
// 	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", d.Price, d.Code)
// 	db.Table("products").Where("id = ?", 1).Update("price", d.Price)
// }

// func ReadHandler(w http.ResponseWriter, r *http.Request) {
// 	var product Product
// 	db := Db_connection()
// 	fmt.Println("read")
// 	if r.Method != "POST" {
// 		return
// 	}
// 	if r.URL.Path != "/read" {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	var d Product

// 	decoder.Decode(&d)
// 	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", d.Price, d.Code)
// 	db.Where("code = ?", d.Code).First(&product)

// 	//read1 := db.First(&product, 1) // find product with id 1
// 	fmt.Println("data: ", product.Code, product.Price)
// }

// Migrate the schema
// db.AutoMigrate(&crud.Product{})

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	fmt.Println(r.Form) // print information on server side.
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello ashish!") // write data to response
// }

//curl -H "Origin: http://example.com" \
// originsOk, headersOk, methodsOk

// func login(w http.ResponseWriter, r *http.Request) {
// 	//w.Header().Set("Access-Control-Allow-Origin", "*")
// 	//r.ParseForm()
// 	user := User{}
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		panic(err)
// 	}
// 	userJson, err := json.Marshal(user)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(userJson)
// 	defer r.Form.Close()
// 	log.Println(user.Username)
// 	log.Println(user.Password)
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()

// 	if err != nil {
// 		panic(err)
// 	}

// 	decoder := schema.NewDecoder()
// 	// r.PostForm is a map of our POST form values
// 	err = decoder.Decode(person, r.PostForm)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Do something with person.Name or person.Phone
// }

// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
// originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
// //corsObj := handlers.AllowedOrigins([]string{"*"})
