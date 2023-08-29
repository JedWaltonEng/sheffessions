package store

import (
	"database/sql"
	"log"
	"os"
)

type Storer interface {
	SaveConfession(content, source string) (int64, error)
	RandomConfession() (*Confession, error)
	DeleteConfessionByID(id int64) error
}

type DBStore struct {
	DB *sql.DB
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

func (store *DBStore) SaveConfession(content, source string) (int64, error) {
	result, err := store.DB.Exec("INSERT INTO confessions (confession_text, source_of_confession) VALUES ($1, $2)", content, source)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

type Confession struct {
	ID                 int64
	ConfessionText     string
	DateOfConfession   string
	SourceOfConfession string
}

func (store *DBStore) RandomConfession() (*Confession, error) {
	query := `
        SELECT * FROM (
            SELECT * 
            FROM confessions 
            WHERE source_of_confession = 'sheffessions' 
            ORDER BY date_of_confession DESC 
            LIMIT 5
        ) AS recent_confessions
        ORDER BY RANDOM()
        LIMIT 1;
    `

	confession := &Confession{}
	err := store.DB.QueryRow(query).Scan(&confession.ID, &confession.ConfessionText, &confession.DateOfConfession, &confession.SourceOfConfession)

	if err != nil {
		return nil, err
	}
	return confession, nil
}

func (store *DBStore) DeleteConfessionByID(id int64) error {
	_, err := store.DB.Exec("DELETE FROM confessions WHERE id = $1", id)
	return err
}

func InitDB() *DBStore {
	connectionString := os.Getenv("POSTGRESQL_URL")
	if connectionString == "" {
		log.Fatal("POSTGRESQL_URL environment variable is not set")
	}
	return NewDBStore(connectionString)
}
