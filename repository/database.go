package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Arun@1435 dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
		
	}
	return db, nil
}
