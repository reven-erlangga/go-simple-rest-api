package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reven-erlangga/go-simple-rest-api/initializers"
	"github.com/reven-erlangga/go-simple-rest-api/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var request struct {
		Email string
		Password string
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{
		Email: request.Email,
		Password: string(hash),
	}
	result := initializers.DB.Create(user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success to create user",
	})
}