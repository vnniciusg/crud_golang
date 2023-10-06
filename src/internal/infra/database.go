package infra

import (
	"database/sql"
	"log"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Println("Error on connect database: ", err)
	}

	return &Database{
		DB: db,
	}
}
