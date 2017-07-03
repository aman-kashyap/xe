package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Env struct {
	db *gorm.DB
}

//var mob []Redmi

type Redmi struct {
	gorm.Model

	ModelID string //`gorm:"primary_key"`
	Name    string `gorm:"not null;unique"`
	Price   int

	//Model     gorm.Model `gorm:"embedded"`
}

var red []Redmi

//Create
func (env *Env) createHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var redmi Redmi
	_ = json.NewDecoder(r.Body).Decode(&redmi)
	redmi.ModelID = params["id"]
	red = append(red, redmi)
	json.NewEncoder(w).Encode(red)
	env.db.Create(&Redmi{Name: "Redmi4A", Price: 5999})
	//params := mux.Vars(r)
	//var redmi Redmi
	/*for _, item := range red {
		if item.ModelID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	//_ = json.NewDecoder(r.Body).Decode(&redmi)

	json.NewEncoder(w).Encode(&Redmi{})*/

}

/*//Read
func (env *Env) readHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

}

//Update
func (env *Env) updateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

}

//Delete
func (env *Env) deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.Delete(&Redmi{})
	//c.JSON(200, mux.H{"id #" + id: "deleted"})
	decoder := json.NewDecoder(r.Body)

}*/
/*func init() {
	//var err error
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}*/

func main() {
	var err error
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	env := &Env{db: db}
	db.AutoMigrate(&Redmi{})
	red = append(red, Redmi{Name: "Redmi4A", Price: 5999})
	//router connection
	r := mux.NewRouter()

	r.HandleFunc("/create", env.createHandler).Methods("POST")
	//r.HandleFunc("/read", env.readHandler).Methods("GET")
	//r.HandleFunc("/update", env.updateHandler).Methods("PUT")
	//r.HandleFunc("/delete", env.deleteHandler).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7890", nil))
}
