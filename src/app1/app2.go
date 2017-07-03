package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	//"reflect"
	"log"
	"net/http"
)

//type Password []byte

type Person struct {
	Name  string
	Phone string
	//Pass  Password
}

//type Phone struct {
//	Label  string
//	Number string
//}

/*<form>
    <input type="text" name="Name">
    <input type="text" name="Phone.Label">
    <input type="text" name="Phone.Number">
</form>*/

/*func AppsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm(){
	if err != nil{
		//fmt.Println(err)
	}
}
	//decoder := schema.NewDecoder()
	err:= decoder.Decode(person, r.PostForm)
	if err != nil{

	}

}*/

var decoder = schema.NewDecoder()

func main() {
	values := map[string][]string{
		"Name":  {"Maga78n"},
		"Phone": {"7298-24-5336"},
		//"Pass":  {"010101010101010"},
	}

	person := new(Person)
	decoder.Decode(person, values)
	//decoder.RegisterConverter(Password(""), convertByteSlice)
	fmt.Printf("%v\n", values)
	/*r := mux.NewRouter()
	r.HandleFunc("/", AppsHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5010", nil))*/

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":5010", r))
}

/*func convertByteSlice(value string) reflect.Value {
	panic("If 'type Password' is []byte, convertByteSlice is not called")
	fmt.Printf("%s", value)
	x := Password(value)
	v := reflect.ValueOf(&x)

	return v
}*/
func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "%v\n", values)
}
