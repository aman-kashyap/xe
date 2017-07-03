//This program prints Url as output
package main

import (
	//"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	//r.HandleFunc("/v1/login/", ApiHandler)
	r.HandleFunc("/v1/{category}", ApiHandler) //.Methods("GET").Schemes("http")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":5000", r))

}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello there!!\n"))
	fmt.Fprintf(w, "hello %s", r.URL.Path[4:])
}
