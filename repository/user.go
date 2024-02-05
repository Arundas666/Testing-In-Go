package repository

import (
	"context"
	"errors"
	"test/domain"
	"test/entity"
	"test/repository/interfaces"

	"github.com/jmoiron/sqlx"
)

var err error

type UserRepo struct {
	Db *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) interfaces.UserRepo {
	return &UserRepo{DB}
}

func (ud *UserRepo) GetByName(ctx context.Context, user entity.CreateUserInput) error {
	query := "SELECT COUNT(*) FROM users WHERE firstname = $1"
	var count int
	err = ud.Db.QueryRow(query, user.FirstName).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}
	return nil
}

func (ud *UserRepo) Create(ctx context.Context, input entity.CreateUserInput) (entity.User, error) {
	result, err := ud.Db.ExecContext(ctx,
		`INSERT INTO users (email, first_name, last_name, gender) VALUES ($1, $2, $3, $4);`,
		input.Email,
		input.FirstName,
		input.LastName,
		input.Gender)
	if err != nil {
		return entity.User{}, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:        userID,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Gender:    input.Gender,
	}, nil
}
var Db *sqlx.DB
func Login(user domain.User) error {
	query := "SELECT COUNT(*) FROM users WHERE firstname = $1 AND password = $2"
	var count int
	err = Db.QueryRow(query, user.FirstName, user.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("invalid username or password")
	}
	return nil
}

func (ud *UserRepo) DeleteUser(user domain.User) error {
	query := "DELETE FROM users WHERE firstname = $1"
	_, err := ud.Db.Exec(query, user.FirstName)
	if err != nil {
		return err
	}
	return nil
}
