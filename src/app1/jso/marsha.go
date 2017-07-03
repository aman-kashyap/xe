// Encoding using json
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Load struct {
	Stuff Data
}
type Data struct {
	Fruit Fruits
	//Veg Vegetables
}
type Fruits map[string]int

//type Vegetables map[string]int

func Handler(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5001", nil))
}

func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Black Grapes"] = 250
	fruits["Orangess"] = 25

	d := Data{fruits}
	l := Load{d}

	return json.MarshalIndent(l, "", "  ")

}
