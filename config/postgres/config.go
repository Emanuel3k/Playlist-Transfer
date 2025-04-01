package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var (
	DATABASE_CONNECTION_PATH = "DATABASE_CONNECTION_PATH"
)

func Config() error {
	dbConnPath := os.Getenv(DATABASE_CONNECTION_PATH)

	db, err := sql.Open("postgres", dbConnPath)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
