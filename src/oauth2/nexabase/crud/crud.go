package crud

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

type Product struct {
	// gorm.Model
	Code  string
	Price uint
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	db := Db_connection()
	fmt.Println("path2", r.URL.Path)
	if r.Method != "POST" {
		return
	}
	if r.URL.Path != "/create" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var d Product

	decoder.Decode(&d)
	db.Create(&Product{Code: d.Code, Price: d.Price})
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", d.Price, d.Code)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	db := Db_connection()
	fmt.Println("delete")
	if r.Method != "POST" {
		return
	}
	if r.URL.Path != "/delete" {
		return
	}

	// Delete - delete product
	db.Delete(&product)
}
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	//var product Product
	db := Db_connection()
	fmt.Println("update")
	if r.Method != "POST" {
		return
	}
	if r.URL.Path != "/update" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var d Product

	decoder.Decode(&d)
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", d.Price, d.Code)
	db.Table("products").Where("id = ?", 1).Update("price", d.Price)
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	db := Db_connection()
	fmt.Println("read")
	if r.Method != "POST" {
		return
	}
	if r.URL.Path != "/read" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var d Product

	decoder.Decode(&d)
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", d.Price, d.Code)
	db.Where("code = ?", d.Code).First(&product)

	//read1 := db.First(&product, 1) // find product with id 1
	fmt.Println("data: ", product.Code, product.Price)
}

func Db_connection() (db *gorm.DB) {

	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable password=postgres")
	check(err)
	db.LogMode(false)
	return db
}
