package main

import (
	"fmt"
	"log"
	"test/handlers"
	"test/repository"
	"test/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewRegisterUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	router := gin.Default()
	fmt.Println("HEyy")
	router.POST("/usersignup", userHandler.Register)
	router.GET("/userlogin", userHandler.Login)

	router.Run(":8000")

}
