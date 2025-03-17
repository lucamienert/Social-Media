package controllers

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment added"})
}
