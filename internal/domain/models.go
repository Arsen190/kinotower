package domain

import "time"

type Gender struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Country struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	FilmCount int    `json:"filmCount,omitempty"`
}

type Category struct {
	ID             int       `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	ParentCategory *Category `json:"parentCategory,omitempty"`
	FilmCount      int       `json:"filmCount,omitempty"`
}

type User struct {
	ID        int       `json:"id" db:"id"`
	Fio       string    `json:"fio" db:"fio"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Birthday  string    `json:"birthday" db:"birthday"`
	GenderID  int       `json:"gender_id" db:"gender_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Film struct {
	ID            int        `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Duration      int        `json:"duration" db:"duration"`
	YearOfIssue   int        `json:"year_of_issue" db:"year_of_issue"`
	Age           int        `json:"age" db:"age"`
	LinkImg       *string    `json:"link_img" db:"link_img"`
	LinkKinopoisk *string    `json:"link_kinopoisk" db:"link_kinopoisk"`
	LinkVideo     string     `json:"link_video" db:"link_video"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	Country       Country    `json:"country"`
	Categories    []Category `json:"categories"`
	RatingAvg     float64    `json:"ratingAvg"`
	ReviewCount   int        `json:"reviewCount"`
}

type FilmListResponse struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int    `json:"total"`
	Films []Film `json:"films"`
}

type Review struct {
	ID         int       `json:"id" db:"id"`
	User       User      `json:"user"`
	Film       *Film     `json:"film,omitempty"` // Используем для списка отзывов пользователя
	Message    string    `json:"message" db:"message"`
	IsApproved bool      `json:"is_approved" db:"is_approved"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type Rating struct {
	ID        int       `json:"id" db:"id"`
	User      User      `json:"user"`
	Film      Film      `json:"film"`
	Ball      int       `json:"ball" db:"ball"` // В ТЗ поле называется ball (балл)
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}