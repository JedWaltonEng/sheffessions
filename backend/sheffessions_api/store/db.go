package store

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	connectionString := os.Getenv("POSTGRESQL_URL")
	if connectionString == "" {
		log.Fatal("POSTGRESQL_URL environment variable is not set")
	}
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping() // This will check if the connection is successful
	if err != nil {
		log.Fatal(err)
	}
}
