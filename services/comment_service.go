package services

import (
	"lucamienert/twitter-clone/models"
	"lucamienert/twitter-clone/repository"
)

type CommentService struct {
	repo *repository.CommentRepository
}

func NewCommentService(repo *repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) AddComment(comment *models.Comment) error {
	return s.repo.AddComment(comment)
}

func (s *CommentService) GetComments(postID uint) ([]models.Comment, error) {
	return s.repo.GetCommentsByPostID(postID)
}
