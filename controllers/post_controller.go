package controllers

import (
	"lucamienert/twitter-clone/models"
	"lucamienert/twitter-clone/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{service: service}
}

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
func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := pc.service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created"})
}

// GetPosts godoc
// @Summary      Get all posts
// @Description  Fetch all posts from the database
// @Tags         Posts
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  map[string]string
// @Router       /posts [get]
func (pc *PostController) GetPosts(c *gin.Context) {
	posts, err := pc.service.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
