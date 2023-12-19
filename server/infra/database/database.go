package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

var db *sql.DB

func NewDatabase() (*Database, error) {
	if db == nil {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
		return &Database{db: db}, nil
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
