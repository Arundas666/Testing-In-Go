package usecase

import (
	"context"
	"test/entity"
	"test/repository/interfaces"
)

type RegisterUserUseCase interface {
	Execute(ctx context.Context, input entity.CreateUserInput) (entity.User, error)
}

type registerUserUseCase struct {
	userRepo interfaces.UserRepo
}

func NewRegisterUserUseCase(
	userRepo interfaces.UserRepo,
) RegisterUserUseCase {
	return &registerUserUseCase{userRepo: userRepo}
}

func (w *registerUserUseCase) Execute(ctx context.Context, input entity.CreateUserInput) (entity.User, error) {

	user, err := w.userRepo.Create(ctx, input)
	if err != nil {
		return entity.User{}, err
	}

	// Do something
	// ...
	// ...

	return user, nil
}
