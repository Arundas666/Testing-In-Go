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
	query := "SELECT COUNT(*) FROM users WHERE username = $1"
	var count int
	err = repo.db.QueryRow(query, user.UserName).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}
	
	return nil
}

func CreateUser(user domain.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = repo.db.Exec(query, user.UserName, user.Password)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func Login(user domain.User) error {
	query := "SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2"
	var count int
	err = repo.db.QueryRow(query, user.UserName, user.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("invalid username or password")
	}

	return nil
}

func DeleteUser(user domain.User) error {
	query := "DELETE FROM users WHERE username = $1"
	_, err := repo.db.Exec(query, user.UserName)
	if err != nil {
		return err
	}
	return nil
}
