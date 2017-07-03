//a simple web server 
/*package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"time"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!"))
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)

	log.Fatal(http.ListenAndServe(":8060", r))

}*/
/*
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
*/