package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

// Movies holds a movie data
type Movies struct {
	Name       string   `bson:"name"`
	Year       string   `bson:"year"`
	Directors  []string `bson:"directors"`
	Writers    []string `bson:"writers"`
	BoxOffices `bson:"boxOffice"`
}

// BoxOffice is nested in Movie
type BoxOffices struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

func main() {
	// load .env file
	loadEnvErr := godotenv.Load()

	if loadEnvErr != nil {
		log.Fatal("Unable to load .env", loadEnvErr)
	}

	dbUri := os.Getenv("MONGODB_URI")
	fmt.Println("dbUri: ", dbUri)
	session, err := mgo.Dial(dbUri)
	c := session.DB("testDb").C("movies")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Create a movie
	darkNight := &Movies{
		Name:      "The Dark Night",
		Year:      "2008",
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffices: BoxOffices{
			Budget: 1850000000,
			Gross:  533316061,
		},
	}

	// Insert into MongoDB
	err = c.Insert(darkNight)
	if err != nil {
		log.Fatal(err)
	}

	// Now query the movie back
	result := Movies{}

	// bson.M is used for nested fields
	err = c.Find(bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie: ", result.Name)
}
