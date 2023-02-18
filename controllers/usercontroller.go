package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/reven-erlangga/go-simple-rest-api/initializers"
	"github.com/reven-erlangga/go-simple-rest-api/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var userRequest models.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{
		Email: userRequest.Email,
		Password: string(hash),
	}
	
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success to create user",
		"user": userRequest,
	})
}

func Login(c *gin.Context)  {
	var userRequest struct {
		Email string
		Password string
	}

	if c.Bind(&userRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", userRequest.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})

		return
	}

	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	// c.JSON(http.StatusOK, gin.H{
		
	// })
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}