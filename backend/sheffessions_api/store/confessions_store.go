package store

type StorerConfessions interface {
	SaveConfession(content, source string) (int64, error)
	RandomConfession() (*Confession, error)
	DeleteConfessionByID(id int64) error
}

type Confession struct {
	ID                 int64
	ConfessionText     string
	DateOfConfession   string
	SourceOfConfession string
}

func (store *Store) SaveConfession(content, source string) (int64, error) {

	result, err := store.db.Exec("INSERT INTO confessions (confession_text, source_of_confession) VALUES ($1, $2)", content, source)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (store *Store) RandomConfession() (*Confession, error) {
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
	err := store.db.QueryRow(query).Scan(&confession.ID, &confession.ConfessionText, &confession.DateOfConfession, &confession.SourceOfConfession)

	if err != nil {
		return nil, err
	}
	return confession, nil
}

func (store *Store) DeleteConfessionByID(id int64) error {
	_, err := store.db.Exec("DELETE FROM confessions WHERE id = $1", id)
	return err
}
