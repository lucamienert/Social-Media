package controllers

import (
	"net/http"
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/middleware"
	"lucamienert/twitter-clone/models"

	"github.com/gin-gonic/gin"
)

// Register a new user
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login and get JWT
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
