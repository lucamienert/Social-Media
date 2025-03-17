package repository

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/models"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) Create(post *models.Post) error {
	result := config.DB.Create(post)
	return result.Error
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	result := config.DB.Find(&posts)
	return posts, result.Error
}
