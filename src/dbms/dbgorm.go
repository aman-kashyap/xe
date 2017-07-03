package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	//"time"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Don{})
	//db.AutoMigrate(&User{}, &Product{}, &Order{})

	// Add table suffix when create tables
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	// Check model `User`'s table exists or not
	//db.HasTable(&User{})

	// Check table `users` exists or not
	//db.HasTable("users")

	type Don struct {
		gorm.Model
		Num int `gorm:"AUTO_INCREMENT" gorm:"primary_key"`
		//Birthday time.Time
		Age  int
		Name string `gorm:"size:255"` // Default size for string is 255, reset it with this tag

	}

	info := Don{Name: "Manoj", Age: 16}
	rinfo := Don{Name: "Rohan", Age: 10}
	pinfo := Don{Name: "Saroj", Age: 18}
	qinfo := Don{Name: "Magan", Age: 12}

	db.NewRecord(info) // => returns `true` as primary key is blank

	db.Create(&info)
	db.Create(&rinfo)
	db.Create(&pinfo)
	db.Create(&qinfo)

	db.NewRecord(info) // => return `false` after `info` created

	db.CreateTable(&info)
	db.Delete(&info)
}
