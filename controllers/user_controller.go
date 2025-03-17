package controllers

import (
	"lucamienert/twitter-clone/models"
	"lucamienert/twitter-clone/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body  map[string]string  true  "User Data"
// @Success      201  {object}  map[string]string
// @Router       /register [post]
func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := uc.userService.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
func (uc *UserController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := uc.userService.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

// Register godoc
// @Summary      Deletes a user
// @Description  Deletes a user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body  map[string]string  true  "User Data"
// @Success      201  {object}  map[string]string
// @Router       /delete [post]
func (uc *UserController) DeleteUser(c *gin.Context) {
	email := c.Param("email")
	user, err := uc.userService.Login(email, "") // Mock authentication
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := uc.userService.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
