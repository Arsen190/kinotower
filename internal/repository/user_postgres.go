package repository

import (
	"github.com/jmoiron/sqlx"
	"kinotower/internal/domain"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// Создание пользователя
func (r *UserPostgres) Create(user domain.User) (int, error) {
	var id int
	query := `INSERT INTO users (fio, email, password, birthday, gender_id) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	
	row := r.db.QueryRow(query, user.Fio, user.Email, user.Password, user.Birthday, user.GenderID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// Поиск по Email (для входа)
func (r *UserPostgres) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, fio, email, password FROM users WHERE email=$1 AND deleted_at IS NULL"
	err := r.db.Get(&user, query, email)
	return user, err
}

// ПОЛУЧЕНИЕ ПО ID (Этого метода, скорее всего, не хватало)
func (r *UserPostgres) GetByID(id int) (domain.User, error) {
	var user domain.User
	query := "SELECT id, fio, email, birthday, gender_id FROM users WHERE id=$1 AND deleted_at IS NULL"
	err := r.db.Get(&user, query, id)
	return user, err
}

// УДАЛЕНИЕ (Этого метода точно не хватало)
func (r *UserPostgres) Delete(id int) error {
	// Делаем Soft Delete (помечаем дату удаления), как в ТЗ
	query := "UPDATE users SET deleted_at = NOW() WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}