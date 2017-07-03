package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//var db *gorm.DB

type Item struct {
	gorm.Model
	Name  string
	Price int
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Item{})
	fmt.Println(CreateItem(&db))
}

func CreateItem(i *Item) {
	var items []Item = []Item{
		Item{Name: "Redmi4A", Price: 5998},
		Item{Name: "Redmi Note 4(2GB)", Price: 9999},
		Item{Name: "Redmi Note 4(3GB)", Price: 10999},
	}

	for _, item := range items {
		db.Create(&item)
	}
}
