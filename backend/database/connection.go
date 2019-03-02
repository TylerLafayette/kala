package database

import (
	"database/sql"
	"log"

	// Import the postgres driver
	_ "github.com/lib/pq"
)

// Database stores a single database object.
type Database struct {
	Conn *sql.DB
}

// DB stores a global database object
var DB Database

// Connect initializes a connection to our database.
func Connect() {
	connString := "user=tyler dbname=tyler sslmode=disable"
	var err error
	DB.Conn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
}
