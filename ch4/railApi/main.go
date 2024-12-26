package main

import (
	"database/sql"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
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
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := dbutils.TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
