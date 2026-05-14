package repository

import (
	"github.com/jmoiron/sqlx"
	"kinotower/internal/domain"
)

type User interface {
	Create(user domain.User) (int, error)
	GetByEmail(email string) (domain.User, error)
	GetByID(id int) (domain.User, error)
	Delete(id int) error
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
	CheckExists(userId, filmId int) (bool, error)
}

type Repository struct {
	User
	Film
	Review
	Rating
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:   NewUserPostgres(db),
		Film:   NewFilmPostgres(db),
		Review: NewReviewPostgres(db), // Подключили
		Rating: NewRatingPostgres(db), // Подключили
	}
}