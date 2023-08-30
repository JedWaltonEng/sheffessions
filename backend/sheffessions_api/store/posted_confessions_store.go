package store

import (
	"database/sql"
	"time"
)

type StorerPostedConfessions interface {
	MarkConfessionAsPublished(confessionID int64) error
	IsConfessionPublished(confessionID int64) (bool, error)
}

type PublishedConfession struct {
	ID           int64     `db:"id"`
	ConfessionID int64     `db:"confession_id"`
	DatePosted   time.Time `db:"date_posted"`
}

func (store *Store) MarkConfessionAsPublished(confessionID int64) (int64, error) {
	result, err := store.db.Exec("INSERT INTO published_confessions (confession_id) VALUES ($1)", confessionID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (store *Store) IsConfessionPublished(confessionID int64) (bool, error) {
	var exists bool
	query := `SELECT exists(SELECT 1 FROM published_confessions WHERE confession_id=$1)`
	err := store.db.QueryRow(query, confessionID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
