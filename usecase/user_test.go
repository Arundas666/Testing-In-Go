package usecase

import (
	"context"
	"fmt"
	"reflect"
	"test/entity"
	mock "test/mock/mockRepo"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_registerUserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockUserRepo(ctrl)
	fmt.Println(c)

	type args struct {
		ctx   context.Context
		input entity.CreateUserInput
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *mock.MockUserRepo)
		want       entity.User
		wantErr    bool
	}{
		{
			name: "success creating new user",
			args: args{
				ctx: context.TODO(),
				input: entity.CreateUserInput{
					Email:     "john.doe@example.com",
					FirstName: "John",
					LastName:  "Doe",
					Gender:    "MALE",
				},
			},
			beforeTest: func(userRepo *mock.MockUserRepo) {
				userRepo.EXPECT().
					Create(
						context.TODO(),
						entity.CreateUserInput{
							Email:     "john.doe@example.com",
							FirstName: "John",
							LastName:  "Doe",
							Gender:    "MALE",
						},
					).
					Return(
						entity.User{
							ID:        1,
							Email:     "john.doe@example.com",
							FirstName: "John",
							LastName:  "Doe",
							Gender:    "MALE",
						},
						nil,
					)
			},
			want: entity.User{
				ID:        1,
				Email:     "john.doe@example.com",
				FirstName: "John",
				LastName:  "Doe",
				Gender:    "MALE",
			},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepo(ctrl)

			w := &registerUserUseCase{
				userRepo: mockUserRepo,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockUserRepo)
			}

			got, err := w.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("registerUserUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerUserUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
