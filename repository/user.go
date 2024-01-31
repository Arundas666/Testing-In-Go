package repository

import (
	"database/sql"
	"errors"
	"log"
	"test/domain"
)

var err error

type Repository struct {
	db *sql.DB
}

var repo Repository

func init() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	repo.db = db
}

func GetByName(user domain.User) error {
	query := "SELECT COUNT(*) FROM users WHERE firstname = $1"
	var count int
	err = repo.db.QueryRow(query, user.FirstName).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}
	return nil
}

func CreateUser(user domain.User) error {
	query := "INSERT INTO users (firstname, password) VALUES ($1, $2)"
	_, err = repo.db.Exec(query, user.FirstName, user.Password)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func Login(user domain.User) error {
	query := "SELECT COUNT(*) FROM users WHERE firstname = $1 AND password = $2"
	var count int
	err = repo.db.QueryRow(query, user.FirstName, user.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("invalid username or password")
	}

	return nil
}

func DeleteUser(user domain.User) error {
	query := "DELETE FROM users WHERE firstname = $1"
	_, err := repo.db.Exec(query, user.FirstName)
	if err != nil {
		return err
	}
	return nil
}
