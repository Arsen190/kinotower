package handler

import (
	"encoding/json"
	"kinotower/internal/domain"
	"kinotower/internal/middleware" // Добавили для UserCtx
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) createReview(w http.ResponseWriter, r *http.Request) {
    // Получаем userId из токена
    userId, _ := r.Context().Value(middleware.UserCtx).(int)

	var input struct {
		FilmID  int    `json:"film_id"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	id, err := h.services.Review.Create(domain.Review{Message: input.Message}, userId, input.FilmID)
    if err != nil {
        newErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }
    
	sendJSON(w, http.StatusCreated, map[string]interface{}{"id": id, "status": "success"})
}

func (h *Handler) getFilmReviews(w http.ResponseWriter, r *http.Request) {
	filmIdStr := chi.URLParam(r, "id")
    filmId, _ := strconv.Atoi(filmIdStr)

	reviews, err := h.services.Review.GetByFilmID(filmId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendJSON(w, http.StatusOK, map[string]interface{}{"reviews": reviews})
}