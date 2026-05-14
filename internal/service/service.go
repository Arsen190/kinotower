package service

import (
	"kinotower/internal/domain"
	"kinotower/internal/repository"
)

type Auth interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	Delete(userId int) error
}

type Film interface {
	GetAll(filters map[string]string) ([]domain.Film, error)
	GetByID(id int) (domain.Film, error)
	GetTotalCount(filters map[string]string) (int, error)
}

// ДОБАВИТЬ ЭТИ ИНТЕРФЕЙСЫ:
type Review interface {
	Create(review domain.Review, userId, filmId int) (int, error)
	GetByFilmID(filmId int) ([]domain.Review, error)
}

type Rating interface {
	Create(userId, filmId, ball int) (int, error)
}

type Service struct {
	Auth
	Film
	Review
	Rating
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:   NewAuthService(repos.User),
		Film:   NewFilmService(repos.Film),
		Review: NewReviewService(repos.Review), // Подключили
		Rating: NewRatingService(repos.Rating), // Подключили
	}
}