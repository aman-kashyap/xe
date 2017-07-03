package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"nexadash-backend-go/basemodel"
	"nexadash-backend-go/models"
)

func AppRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		args, _ := r.URL.Query()["limit"]
		fmt.Println(reflect.TypeOf(args))
		fetched_value := basemodel.Fetch(r)
		fmt.Println(json.NewEncoder(w).Encode(fetched_value))
	}
	if r.Method == "POST" {
		post := models.Make_app_name(r)
		fmt.Println(json.NewEncoder(w).Encode(post))
	}
	if r.Method == "PUT" {
		fmt.Println("put request is working")
		update := basemodel.Update(r)
		fmt.Println(json.NewEncoder(w).Encode(update))
	}
}

func GetSpecific(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ID := basemodel.Return_id(r)
		all_records := basemodel.Fetch_by_id(r, ID)
		fmt.Println(json.NewEncoder(w).Encode(all_records))
	}
	if r.Method == "DELETE" {
		ID := basemodel.Return_id(r)
		deleted_record := basemodel.Remove_by_id(r, ID)
		fmt.Println(json.NewEncoder(w).Encode(deleted_record))
	}
}

func main() {
	r := mux.NewRouter()
	db := basemodel.Db_connection()
	fmt.Println(db)
	defer db.Close()
	fmt.Println("database", db)
	// Migrate the schema
	db.AutoMigrate(&basemodel.Apps8{})

	r.HandleFunc("/v1/apps", AppRequestHandler)
	r.HandleFunc("/v1/apps/{id}", GetSpecific)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
