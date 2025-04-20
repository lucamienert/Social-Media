package controllers

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/lucamienert/Social-Media/models"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// GetMe godoc
// @Summary Get current authenticated user
// @Description Get the details of the currently authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "status: success, data: {user: models.UserResponse}"
// @Failure 400 {object} map[string]string "status: fail, message: Bad request"
// @Failure 401 {object} map[string]string "status: fail, message: Unauthorized"
// @Security BearerAuth
// @Router /api/user/me [get]
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
