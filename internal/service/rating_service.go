package service

import (
	"errors"
	"kinotower/internal/repository"
)

type RatingService struct {
	repo repository.Rating
}

func NewRatingService(repo repository.Rating) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) Create(userId, filmId, ball int) (int, error) {
	// 1. Проверяем, существует ли уже оценка
	exists, err := s.repo.CheckExists(userId, filmId)
	if err != nil {
		return 0, err
	}
	if exists {
		// Возвращаем специфичную ошибку для ТЗ
		return 0, errors.New("Score exist")
	}

	// 2. Если нет, создаем
	return s.repo.Create(userId, filmId, ball)
}