package controllers

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost godoc
// @Summary      Create a post
// @Description  Creates a new post (requires authentication)
// @Tags         Posts
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        post  body  map[string]string  true  "Post Data"
// @Success      201  {object}  map[string]string
// @Router       /post [post]
func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post created"})
}

// GetPosts godoc
// @Summary      Get all posts
// @Description  Fetch all posts from the database
// @Tags         Posts
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  map[string]string
// @Router       /posts [get]
func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}
