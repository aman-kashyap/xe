package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type Card struct { //defining the data structure for our virtual flashcards
	Term       string `json:"Term"`
	Definition string `json:"Definition"`
}

func open(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("card.html")
	t.Execute(w, nil)
}

func addcard(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("deck.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	card := new(Card)
	card.Term = r.FormValue("term")
	card.Definition = r.FormValue("definition")

	b, err := json.Marshal(card)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	f.Write(b)
	f.Close()
}

func main() {
	http.HandleFunc("/addcard", addcard)
	http.HandleFunc("/", open)
	http.ListenAndServe(":8080", nil)

}
