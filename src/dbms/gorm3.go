package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

/*type Product struct {
	gorm.Model
	Code  string
	Price uint
}*/

type Information struct {
	gorm.Model
	Name     string
	Age      int `gorm:"size:255"`
	Interest string
	//Emails   []Email
}

/*type Email struct {
	id    int
	Email string `gorm:"type:varchar(100);unique_index"`
}*/

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Information{})

	// Create
	db.Create(&Information{Name: "Amit", Age: 22, Interest: "Node.js"})

	// Read
	var information Information
	db.First(&information, 1)               // find information with id 1
	db.First(&information, "age = ?", "22") // find information with age 20

	// Update - update information's price to 2000
	db.Model(&information).Update("Age", 23)

	// Delete - delete information
	//db.Delete(&information)
}
