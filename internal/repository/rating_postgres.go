package repository

import "github.com/jmoiron/sqlx"

type RatingPostgres struct {
	db *sqlx.DB
}

func NewRatingPostgres(db *sqlx.DB) *RatingPostgres {
	return &RatingPostgres{db: db}
}

func (r *RatingPostgres) Create(userId, filmId, ball int) (int, error) {
	var id int
	query := `INSERT INTO ratings (user_id, film_id, rating) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, userId, filmId, ball).Scan(&id)
	return id, err
}

func (r *RatingPostgres) CheckExists(userId, filmId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM ratings WHERE user_id = $1 AND film_id = $2`
	err := r.db.Get(&count, query, userId, filmId)
	return count > 0, err
}