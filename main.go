package main

import (
	"books/models"
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:@localhost:5432/books?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	server := models.NewServer(db)

	http.ListenAndServe(":8080", server.Router)

}
