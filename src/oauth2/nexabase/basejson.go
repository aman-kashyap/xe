package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Payload struct {
	Stuff Data
}
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}
type Fruits map[string]int
type Vegetables map[string]int

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 10

	vegetables := make(map[string]int)
	vegetables["Carrats"] = 10
	vegetables["Beets"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
