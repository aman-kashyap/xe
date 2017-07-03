package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

var err error

func main() {

	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.Create(Country{})

	r := mux.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":1950", r))
}

type Country struct {
	Id         string `json:"id"`
	State      string `json:"state"`
	Population int    `json:"population"`
}
