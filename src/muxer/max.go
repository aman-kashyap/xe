package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//func HomeHandler(w http.ResponseWriter, r *http.Request) {
//w.Write([]byte("Gorilla!\n"))
//	fmt.Fprintf(w, "hi there, I use product of %s! ", r.URL.Path[1:])
//}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillas! Product"))

}

//func ArticleHandler(w http.ResponseWriter, r *http.Request) {

//w.Write([]byte("Gorilla! Articles"))
//}

//func ProcessPathVariables(w http.ResponseWriter, r *http.Request) {

// break down the variables for easier assignment
//	vars := mux.Vars(r)
//	category := vars["category"]
//name := vars["name"]
//job := vars["job"]
//age := vars["age"]
// w.Write([]byte(fmt.Sprintf("Name is %s \n", name)))
// w.Write([]byte(fmt.Sprintf("Job is %s \n", job)))
// w.Write([]byte(fmt.Sprintf("Age is %s \n", age)))

//	route := mux.CurrentRoute(r)
//	w.Write([]byte(fmt.Sprintf("Route name is %v \n", route.GetName())))

//}

//func (r *Route) URL(pairs ...string) (*url.URL, error)

func main() {
	r := mux.NewRouter()

	/*s := r.Host("{subdomain}.xenon.com").Subrouter()
	s.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(ArticleHandler).Name("article")
	url, err := r.Get("article").URL("subdomain", "aman", "category", "go", "id", "3341")

	if err != nil {
		fmt.Println(err)
	}
	builtURL := url.String()
	fmt.Printf("Url :%v \n", builtURL)*/
	//r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{key}", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
type Girl struct{
	Firstname string json:"firstname,omitempty"
	Lastname string json:"lastname,omitempty"
}
var larki []Girl
func GetGirlEndpoint(w http.ResponseWriter, r *http.Request){
	params :=mux.Vars(r)
	for _, item :=range larki{
		json.NewEncoder(w).Encode(item)
		return
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetLarkiEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func main(){
	r := mux.NewRouter()
	larki=append(larki, Girl{Firstname:"Annu", Lastname"singh"})

}
*/
