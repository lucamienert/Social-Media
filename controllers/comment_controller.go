package controllers

import (
	"lucamienert/twitter-clone/models"
	"lucamienert/twitter-clone/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service *services.CommentService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{service: service}
}

// AddComment godoc
// @Summary      Add a comment
// @Description  Adds a comment to a specific post
// @Tags         Comments
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        post_id  path  int  true  "Post ID"
// @Param        comment  body  models.Comment  true  "Comment Data"
// @Success      201  {object}  map[string]string
// @Router       /post/{post_id}/comment [post]
func (cc *CommentController) AddComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := cc.service.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment added"})
}

// GetComments godoc
// @Summary      Get comments for a post
// @Description  Retrieves all comments for a specific post
// @Tags         Comments
// @Produce      json
// @Param        post_id  path  int  true  "Post ID"
// @Success      200  {array}  models.Comment
// @Router       /post/{post_id}/comments [get]
func (cc *CommentController) GetComments(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	comments, err := cc.service.GetComments(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
