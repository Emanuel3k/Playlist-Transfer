package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var (
	DatabaseConnectionPath = "DATABASE_CONNECTION_PATH"
	conn                   *sql.DB
)

func Config() error {
	dbConnPath := os.Getenv(DatabaseConnectionPath)

	db, err := sql.Open("postgres", dbConnPath)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	conn = db

	return nil
}

func GetDB() *sql.DB {
	return conn
}
