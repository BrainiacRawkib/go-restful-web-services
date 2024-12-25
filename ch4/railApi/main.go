package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"railapi/dbutils"
)

func main() {
	// Connect to Database
	db, err := sql.Open("sqlite3", "./railapi.db")

	if err != nil {
		log.Println("Driver creation failed!: ", err)
	}

	// Create tables
	dbutils.Initialize(db)
}
