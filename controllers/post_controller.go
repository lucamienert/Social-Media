package controllers

import (
	"lucamienert/twitter-clone/models"
	"lucamienert/twitter-clone/services"
	"net/http"
	"strconv"

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
// @Param        post  body  models.Post  true  "Post Data"
// @Success      201  {object}  map[string]string
// @Router       /post [post]
func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := pc.service.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created"})
}

// GetPosts godoc
// @Summary      Get all posts
// @Description  Fetch all posts from the database
// @Tags         Posts
// @Produce      json
// @Success      200  {array}  models.Post
// @Router       /posts [get]
func (pc *PostController) GetPosts(c *gin.Context) {
	posts, err := pc.service.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// LikePost godoc
// @Summary      Like a post
// @Description  Increments the like count of a specific post
// @Tags         Posts
// @Security     BearerAuth
// @Param        id  path  int  true  "Post ID"
// @Success      200  {object}  map[string]string
// @Router       /post/{id}/like [post]
func (pc *PostController) LikePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	if err := pc.service.LikePost(uint(postID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post liked"})
}
