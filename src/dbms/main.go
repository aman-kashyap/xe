package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"time"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Company struct {
	ID        int        `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string     `sql:"size:255;unique;index"`
	Employees []Employee // one-to-many relationship
	Address   Address    // one-to-one relationship
}

type Employee struct {
	FirstName        string    `sql:"size:255;index:name_idx"`
	LastName         string    `sql:"size:255;index:name_idx"`
	SocialSecurityNo string    `sql:"type:varchar(100);unique" gorm:"column:ssn"`
	DateOfBirth      time.Time `sql:"DEFAULT:current_timestamp"`
	Address          *Address  // one-to-one relationship
	Deleted          bool      `sql:"DEFAULT:false"`
}

type Address struct {
	Country  string `gorm:"primary_key"`
	City     string `gorm:"primary_key"`
	PostCode string `gorm:"primary_key"`
	Line1    sql.NullString
	Line2    sql.NullString
}

/*type Product struct {
	gorm.Model
	Code  string
	Price uint
}*/

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=aman sslmode=disable password=postgres") // "postgresql://myapp:dbpass@localhost:15432/myapp"
	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	/*db.AutoMigrate(&User{}, &Product{}, &Order{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	db.HasTable(&User{})*/

	db.SingularTable(true)

	db.AutoMigrate(&company.Address{})
	db.AutoMigrate(&company.Company{})
	db.AutoMigrate(&company.Employee{})

	// Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	db.CreateTable(&company.Address{})
	db.CreateTable(&company.Company{})
	db.CreateTable(&company.Employee{})

	db.Model(&company.Company{}).ModifyColumn("name", "text")
	// drop

	db.DropTable(&company.Address{})
	db.DropTable(&company.Company{})
	db.DropTable(&company.Employee{})

	sampleCompany := company.Company{
		Name: "Google",
		Address: company.Address{
			Country:  "USA",
			City:     "Moutain View",
			PostCode: "1600",
		},
		Employees: []company.Employee{
			company.Employee{
				FirstName:        "John",
				LastName:         "Doe",
				SocialSecurityNo: "00-000-0000",
			},
		},
	}

	db.Create(&sampleCompany)

	db.Where("Name LIKE ?", "%G%").Delete(company.Company{})

	model.Country = "USA"
	db.Save(&sampleCompany)

	db.Table("addresses").Where("Country = ?", "USA").Updates(map[string]interface{}{"Country": "North America"})

	var firstComp company.Company

	// fetch a company by primary key
	db.First(&firstComp, 1)

	// fetch a company by name
	db.Find(&firstComp, "name = ?", "Google")

	// fetch all companies
	var comapnies []company.Company
	db.Find(&companies)

	// fetch all companies that starts with G
	db.Where("name = ?", "%G%").Find(&companies)

	// Read
	/*var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)*/
}
