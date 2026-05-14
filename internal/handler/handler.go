package handler

import (
	"kinotower/internal/middleware"
	"kinotower/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		// Публичные
		r.Post("/auth/signup", h.signUp)
		r.Post("/auth/signin", h.signIn)
		r.Get("/films", h.getFilms)
		r.Get("/films/{id}", h.getFilmById)
		r.Get("/films/{id}/reviews", h.getFilmReviews)
		r.Get("/genders", h.getGenders)

		// Защищенные
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(h.services))
			
			r.Delete("/users", h.deleteUser)
			r.Post("/users/{id}/reviews", h.createReview)
			r.Post("/users/{id}/ratings", h.createRating)
		})
	})

	return r
}

// Оставляем только те, которых еще нет в других файлах
func (h *Handler) getGenders(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"genders": [{"id": 1, "name": "Мужской"}, {"id": 2, "name": "Женский"}]}`))
}