package main

import (
	//"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Information struct {
	gorm.Model
	Name     string
	Age      int `gorm:"size:255"`
	Interest string
}

var (
	db *gorm.DB
)

func init() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func main() {

	r := mux.NewRouter()

	//r.HandleFunc("/", handleMigrate)
	r.HandleFunc("/Information/{id}", handleCreate)
	//r.HandleFunc("/", handleRead)
	//r.HandleFunc("/", handleUpdate)
	//r.HandleFunc("/", handleDelete)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8900", r))
}

/*func handleMigrate(w http.ResponseWriter, r *http.Request) {
	// Migrate the schema
	r.AutoMigrate(&Information{})
}*/

func handleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html.EscapeString(r.URL.Path))
	// Create
	r.Create(&Information{Name: "Amit", Age: 22, Interest: "Node.js"})
}

/*func handleRead(w http.ResponseWriter, r *http.Request) {
	// Read
	var information Information
	r.First(&information, 1)               // find information with id 1
	r.First(&information, "age = ?", "22") // find information with age 20
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	// Update
	r.Model(&information).Update("Age", 22)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	// Delete - delete information
	r.Delete(&information)
}*/
