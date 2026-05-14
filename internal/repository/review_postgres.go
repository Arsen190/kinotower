package repository

import (
	"kinotower/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func NewReviewPostgres(db *sqlx.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}

func (r *ReviewPostgres) Create(review domain.Review, userId, filmId int) (int, error) {
	var id int
	query := `INSERT INTO reviews (user_id, film_id, message) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, userId, filmId, review.Message).Scan(&id)
	return id, err
}

func (r *ReviewPostgres) GetByFilmID(filmId int) ([]domain.Review, error) {
	var reviews []domain.Review
	query := `
		SELECT r.id, r.message, r.created_at, u.id "user.id", u.fio "user.fio"
		FROM reviews r
		JOIN users u ON r.user_id = u.id
		WHERE r.film_id = $1 AND r.is_approved = true
		ORDER BY r.created_at DESC`
	
	rows, err := r.db.Queryx(query, filmId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var rev domain.Review
		if err := rows.Scan(&rev.ID, &rev.Message, &rev.CreatedAt, &rev.User.ID, &rev.User.Fio); err != nil {
			return nil, err
		}
		reviews = append(reviews, rev)
	}
	return reviews, nil
}