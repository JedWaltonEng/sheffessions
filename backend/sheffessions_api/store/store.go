package store

import (
	"database/sql"
	"log"
	"os"
)

type StorerConfessions interface {
	SaveConfession(content, source string) (int64, error)
	RandomConfession() (*Confession, error)
	DeleteConfessionByID(id int64) error
}

type DBStore struct {
	DB *sql.DB
}

type Confession struct {
	ID                 int64
	ConfessionText     string
	DateOfConfession   string
	SourceOfConfession string
}

func NewDBStore(connectionString string) *DBStore {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping() // This will check if the connection is successful
	if err != nil {
		log.Fatal(err)
	}

	return &DBStore{DB: db}
}

func (store *DBStore) Close() error {
	return store.DB.Close()
}

func InitDB() *DBStore {
	connectionString := os.Getenv("POSTGRESQL_URL")
	if connectionString == "" {
		log.Fatal("POSTGRESQL_URL environment variable is not set")
	}
	return NewDBStore(connectionString)
}
