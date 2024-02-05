package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=Arun@1435 dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
