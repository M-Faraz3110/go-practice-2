package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDBContext(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	_, err = db.Exec(migration)
	if err != nil {
		return nil, err
	}
	return db, nil
}
