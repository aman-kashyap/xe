package main

import "fmt"
import "oauth2/nexabase/crud"
import "net/http"

func main() {
	db := crud.Db_connection()
	defer db.Close()
	fmt.Println("database", db)

	// Migrate the schema
	db.AutoMigrate(&crud.Product{})

	http.HandleFunc("/create", crud.CreateHandler)
	http.HandleFunc("/delete", crud.DeleteHandler)
	http.HandleFunc("/update", crud.UpdateHandler)
	http.HandleFunc("/read", crud.ReadHandler)
	http.ListenAndServe(":8000", nil)
}
