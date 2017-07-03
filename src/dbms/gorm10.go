package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var people []Person

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})
	people = append(people, Person{ID: "1", FirstName: "Ramesh", LastName: "Kumar"})
	people = append(people, Person{ID: "2", FirstName: "Santosh", LastName: "Kaushal"})
	//db.Create(&p1)
	//fmt.Println(p1.FirstName)

	r := mux.NewRouter()
	r.HandleFunc("/people/", GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/people", CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7890", r))
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	ID := params["id"]
	var person Person
	people = append(people, person)
	d := db.Where("id = ?", ID).Delete(&person)
	fmt.Println(d)
	json.NewEncoder(w).Encode(people)
}

func UpdatePerson(w http.ResponseWriter, req *http.Request) {

	var person Person
	params := mux.Vars(req)
	ID := params["id"]
	people = append(people, person)

	if err := db.Where("id = ?", ID).First(&person).Error; err != nil {
		fmt.Println(err)
	}
	json.NewDecoder(req.Body).Decode(&person)

	db.Save(&person)

}

func CreatePerson(w http.ResponseWriter, req *http.Request) {

	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	people = append(people, person)
	db.Create(&person)
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	ID := params["id"]
	var person Person
	if err := db.Where("id = ?", ID).First(&person).Error; err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(&Person{})
	}
}
func GetPeople(w http.ResponseWriter, req *http.Request) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(people)
	}
}

/*
   //Parse url parameters passed, then parse the response packet for the POST body (request body)
   // attention: If you do not call ParseForm method, the following data can not be obtained form
   // logic part of log in
   // setting router rule
   // setting listening port

*/
