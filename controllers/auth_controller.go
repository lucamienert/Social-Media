package controllers

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/middleware"
	"lucamienert/twitter-clone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body  map[string]string  true  "User Data"
// @Success      201  {object}  map[string]string
// @Router       /register [post]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticate user credentials
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  map[string]string  true  "Login Credentials"
// @Success      200  {object}  map[string]string
// @Router       /login [post]
func Login(c *gin.Context) {
	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Where("username = ?", input.Username).First(&user)

	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := middleware.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
