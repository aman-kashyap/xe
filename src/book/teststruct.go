package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Data struct {
	Name  string
	Hours int
}

func save(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	hours, err := strconv.Atoi(r.FormValue("Hours"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := &Data{name, hours}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	f, err := os.Open("somefile.json")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	f.Write(b)
	f.Close()
}

// func init() {
// 	http.HandleFunc("/save", save)
// }
func main() {
	fmt.Println("Hello, playground")
	http.HandleFunc("/save", save)
}

// package main

// import (
// 	"fmt"
// )

// type User struct {
// 	Emails   []Email   `json:"emails"`
// 	Accounts []Account `json:"accounts"`
// }

// type Email struct {
// 	ID    int    `json:"id"`
// 	Email string `json:"email"`
// }

// type Account struct {
// 	Github   bool `json:"github"`
// 	Gmail    bool `json:"gmail"`
// 	Twitter  bool `json:"twitter"`
// 	Nexamail bool `json:"nexamail"`
// }

// type Data struct {
// 	Input User `json:"input"`
// }

// func main() {
// 	e := Email{Email: "aman@xenondigilabs.com"}
// 	a := Account{}
// 	usr := User{Emails: []Email{e}, Accounts: []Account{a}}
// 	x := Data{Input: usr}
// 	fmt.Printf("x is '%+v' and '%+v' \n", x.Input.Emails[0], x.Input.Accounts[0])
// }
