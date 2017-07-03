package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello Gophers!!!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run", HandleMain)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
