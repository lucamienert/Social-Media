package controllers

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post created"})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}
