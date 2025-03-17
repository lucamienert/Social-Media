package services

import (
	"lucamienert/twitter-clone/models"
	"lucamienert/twitter-clone/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.Create(post)
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	return s.repo.GetAll()
}
