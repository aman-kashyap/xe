package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Env struct {
	db *gorm.DB
}

type Person struct {
	gorm.Model
	Num       string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	gorm.Model
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func (env *Env) GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range people {
		if item.Num == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func (env *Env) GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)

}

func (env *Env) CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Num = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func (env *Env) DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for index, item := range people {
		if item.Num == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})
	db.Create(&Person{Num: "1", Firstname: "rau", Lastname: "kaka", Address: &Address{City: "jammu", State: "jammu and kashmir"}})
	db.Create(&Person{Num: "2", Firstname: "sonia", Lastname: "sharma", Address: &Address{City: "jammu", State: "jammu and kashmir"}})

	db.Delete(&Person{}, "firstname LIKE ?", "%ra%")

	env := &Env{db: db}

	people = append(people, Person{Num: "1", Firstname: "rau", Lastname: "kaka", Address: &Address{City: "jammu", State: "jammu and kashmir"}})
	people = append(people, Person{Num: "2", Firstname: "sonia", Lastname: "sharma"})

	r := mux.NewRouter()
	r.HandleFunc("/people", env.GetPeopleEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}", env.GetPersonEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}", env.CreatePersonEndpoint).Methods("POST")
	r.HandleFunc("/people/{id}", env.DeletePersonEndpoint).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":12345", r))
}
