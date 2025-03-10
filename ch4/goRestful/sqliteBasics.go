package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Book is a placeholder for book
type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Println(err)
	}

	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64)NULL)")
	if err != nil {
		log.Println("Error creating table `books`!")
	} else {
		log.Println("Successfully created table `books`!")
	}
	statement.Exec()

	// Create
	statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	statement.Exec("A Table of Two Cities", "Charles Dickens", 140430547)
	log.Println("Entry successful...")

	// Read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, book: %s, Author: %s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	// Update
	statement, _ = db.Prepare("update books set name=? where id=?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated the book...")

	// Delete
	statement, _ = db.Prepare("delete from books where id=?")
	statement.Exec(1)

	log.Println("Book deleted successfully...")
}
