package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "aman"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	   INSERT INTO users (age, email, first_name, last_name)
	   VALUES ($1, $2, $3, $4)
	   RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 30, "amankashyap@gmai.com", "aman", "kumar").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New Record ID is:", id)
	fmt.Println("connected to postgresql !!")

}
