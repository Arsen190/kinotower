package service

import (
	"kinotower/internal/domain"
	"kinotower/internal/repository"
)

type FilmService struct {
	repo repository.Film
}

func NewFilmService(repo repository.Film) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) GetAll(filters map[string]string) ([]domain.Film, error) {
	return s.repo.GetAll(filters)
}

func (s *FilmService) GetByID(id int) (domain.Film, error) {
	return s.repo.GetByID(id)
}

func (s *FilmService) GetTotalCount(filters map[string]string) (int, error) {
	return s.repo.GetTotalCount(filters)
}