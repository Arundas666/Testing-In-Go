package tests

import (
	"test/domain"
	"test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	testUser := domain.User{
		FirstName: "Aruna",
		Password:  "Hey",
	}
	if err := repository.GetByName(testUser); err != nil {
		if err.Error() != "username already exists" {
			t.Errorf("Expected error is 'username already exists', got '%s'", err)

		}
	} else {
		t.Errorf("Expected error is 'username already exists', got no error")

	}
}

func TestCreateUser(t *testing.T) {
	testUser := domain.User{
		FirstName: "Arun",
		Password:  "Hey",
	}
	err := repository.CreateUser(testUser)
	assert.NoError(t, err, "expect no error")

	err = repository.DeleteUser(testUser)
	assert.NoError(t, err, "expect no error")

	err = repository.GetByName(testUser)
	assert.NoError(t, err, "expect no error")

}

func TestLogin(t *testing.T) {
	testUser := domain.User{
		FirstName: "Arun",
		Password:  "Hey",
	}
	err := repository.Login(testUser)
	assert.Error(t, err, "expected error")
	assert.EqualError(t, err, "invalid username or password")

}
