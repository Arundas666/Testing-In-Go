package handlers

import (
	"fmt"
	"net/http"
	"test/domain"
	"test/entity"
	"test/repository"
	"test/usecase"

	"github.com/gin-gonic/gin"
)

type userDelivery struct {
	registerUserUseCase usecase.RegisterUserUseCase
}

func NewUserHandler(regUsecase usecase.RegisterUserUseCase) *userDelivery {
	return &userDelivery{
		registerUserUseCase: regUsecase,
	}
}
func (u *userDelivery) Register(c *gin.Context) {
	var createUserInput entity.CreateUserInput

	if err := c.BindJSON(&createUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := u.registerUserUseCase.Execute(c.Request.Context(), createUserInput)
	fmt.Println("Inside Register function")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (uh *userDelivery) Login(c *gin.Context) {
	var input domain.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.Login(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "user loged in succesfuly"})
	}

}
