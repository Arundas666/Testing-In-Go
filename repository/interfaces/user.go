package interfaces

import (
	"context"
	"test/domain"
	"test/entity"
)

type UserRepo interface {
	Create(ctx context.Context, input entity.CreateUserInput) (entity.User, error)
	GetByName(ctx context.Context, input entity.CreateUserInput) error
	DeleteUser(user domain.User) error
}
