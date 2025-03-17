package repository

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/models"
)

type CommentRepository struct{}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{}
}

func (r *CommentRepository) AddComment(comment *models.Comment) error {
	return config.DB.Create(comment).Error
}

func (r *CommentRepository) GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := config.DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
