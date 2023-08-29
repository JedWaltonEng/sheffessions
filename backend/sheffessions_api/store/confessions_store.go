package store

func (store *DBStore) SaveConfession(content, source string) (int64, error) {
	result, err := store.DB.Exec("INSERT INTO confessions (confession_text, source_of_confession) VALUES ($1, $2)", content, source)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
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
