package api

import "database/sql"

func Start() error {

	db, err := setupDatabase("./storage.db")
	if err != nil {
		return err
	}
	defer db.Close()

	
	return nil
}

func setupDatabase(url string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS car (
			id integer primary key autoincrement,
			model text unique not null,
			release_year integer not null,
			description text,
			image blob			
		);
		CREATE INDEX IF NOT EXISTS idx_model ON car(model);`,
	)

	if _, err := stmt.Exec(); err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}