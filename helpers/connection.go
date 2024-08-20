package helpers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	connStr := "root:root@tcp(127.0.0.1:3306)/arifptmx"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
