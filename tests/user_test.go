package tests

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"test/entity"
	"test/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

const GenderFemale = "FEMALE"
const GenderMale = "MALE"
type userRepo struct {
	Db *sqlx.DB
}

func Test_userRepo_Create(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx   context.Context
		input entity.CreateUserInput
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       entity.User
		wantErr    bool
	}{
		{
			name: "fail create user",
			args: args{
				ctx:   context.TODO(),
				input: entity.CreateUserInput{Email: "jane@example.com", FirstName: "Jane", LastName: "Doe", Gender: GenderFemale},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`
						INSERT INTO users (email, first_name, last_name, gender)
						VALUES ($1, $2, $3, $4);`,
					)).
					WithArgs("jane@example.com", "Jane", "Doe", GenderFemale).
					WillReturnError(errors.New("whoops, error"))
			},
			wantErr: true,
		},
		{
			name: "success create user",
			args: args{
				ctx:   context.TODO(),
				input: entity.CreateUserInput{Email: "john@example.com", FirstName: "John", LastName: "Doe", Gender: GenderMale},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`
						INSERT INTO users (email, first_name, last_name, gender)
						VALUES ($1, $2, $3, $4);`,
					)).
					WithArgs("john@example.com", "John", "Doe", GenderMale).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: entity.User{ID: 1, Email: "john@example.com", FirstName: "John", LastName: "Doe", Gender: GenderMale},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			
			u := &repository.UserRepo{
				Db: db,
			}
			// u:=repository.NewUserRepository(db)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.Create(tt.args.ctx, tt.args.input)
			fmt.Println("ERRR",err)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepo.Create() error = %v, wantErr =%v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}

}
