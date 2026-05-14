package service

import (
	"kinotower/internal/domain"
	"kinotower/internal/repository"
)

type ReviewService struct {
	repo repository.Review
}

func NewReviewService(repo repository.Review) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) Create(review domain.Review, userId, filmId int) (int, error) {
	return s.repo.Create(review, userId, filmId)
}

func (s *ReviewService) GetByFilmID(filmId int) ([]domain.Review, error) {
	return s.repo.GetByFilmID(filmId)
}