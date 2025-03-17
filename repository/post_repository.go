package repository

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/models"

	"gorm.io/gorm"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) Create(post *models.Post) error {
	return config.DB.Create(post).Error
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := config.DB.Preload("Comments").Find(&posts).Error
	return posts, err
}

func (r *PostRepository) LikePost(postID uint) error {
	return config.DB.Model(&models.Post{}).Where("id = ?", postID).Update("likes", gorm.Expr("likes + ?", 1)).Error
}
