package repository

import (
	"fmt"
	"kinotower/internal/domain"
	"github.com/jmoiron/sqlx"
)

type FilmPostgres struct {
	db *sqlx.DB
}

func NewFilmPostgres(db *sqlx.DB) *FilmPostgres {
	return &FilmPostgres{db: db}
}

func (r *FilmPostgres) GetAll(filters map[string]string) ([]domain.Film, error) {
	var films []domain.Film

	query := `
		SELECT 
			f.id, f.name, f.duration, f.year_of_issue, f.age, f.link_img, f.link_kinopoisk, f.link_video, f.created_at,
			c.id "country.id", c.name "country.name",
			(SELECT COALESCE(AVG(rating), 0) FROM ratings WHERE film_id = f.id) as rating_avg,
			(SELECT COUNT(*) FROM reviews WHERE film_id = f.id AND is_approved = true) as review_count
		FROM films f
		LEFT JOIN countries c ON f.country_id = c.id
		WHERE f.deleted_at IS NULL`

	if val, ok := filters["search"]; ok && val != "" {
		query += fmt.Sprintf(" AND f.name ILIKE '%%%s%%'", val)
	}
	if val, ok := filters["country"]; ok && val != "0" && val != "" {
		query += fmt.Sprintf(" AND f.country_id = %s", val)
	}

	query += fmt.Sprintf(" ORDER BY %s %s", filters["sortBy"], filters["sortDir"])
	query += fmt.Sprintf(" LIMIT %s OFFSET %s", filters["size"], filters["offset"])

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f domain.Film
		err := rows.Scan(
			&f.ID, &f.Name, &f.Duration, &f.YearOfIssue, &f.Age, &f.LinkImg, &f.LinkKinopoisk, &f.LinkVideo, &f.CreatedAt,
			&f.Country.ID, &f.Country.Name,
			&f.RatingAvg, &f.ReviewCount,
		)
		if err != nil {
			return nil, err
		}
		films = append(films, f)
	}
	return films, nil
}

func (r *FilmPostgres) GetByID(id int) (domain.Film, error) {
	var film domain.Film
	query := `
		SELECT f.id, f.name, f.duration, f.year_of_issue, f.age, f.link_img, f.link_kinopoisk, f.link_video, f.created_at,
		       c.id "country.id", c.name "country.name"
		FROM films f
		LEFT JOIN countries c ON f.country_id = c.id
		WHERE f.id = $1 AND f.deleted_at IS NULL`
	
	row := r.db.QueryRowx(query, id)
	err := row.Scan(
		&film.ID, &film.Name, &film.Duration, &film.YearOfIssue, &film.Age, &film.LinkImg, &film.LinkKinopoisk, &film.LinkVideo, &film.CreatedAt,
		&film.Country.ID, &film.Country.Name,
	)
	return film, err
}

func (r *FilmPostgres) GetTotalCount(filters map[string]string) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM films WHERE deleted_at IS NULL"
	
	if val, ok := filters["search"]; ok && val != "" {
		query += fmt.Sprintf(" AND name ILIKE '%%%s%%'", val)
	}

	err := r.db.Get(&count, query)
	return count, err
}