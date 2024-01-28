package handlers

import (
	"net/http"
	"test/domain"
	"test/repository"

	"github.com/gin-gonic/gin"
)

func Hi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Hi": "Yes it seems good"})
}

func Signup(c *gin.Context) {
	var input domain.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.GetByName(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "creating user failed"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "user created succesfuly"})
	}
}

func Login(c *gin.Context) {
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
