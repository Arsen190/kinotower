package handler

import (
	"encoding/json"
	"kinotower/internal/middleware"
	"net/http"
)

func (h *Handler) createRating(w http.ResponseWriter, r *http.Request) {
	// 1. Получаем ID пользователя из контекста (установлен Middleware)
	userId, ok := r.Context().Value(middleware.UserCtx).(int)
	if !ok {
		newErrorResponse(w, http.StatusUnauthorized, "user not found in context")
		return
	}

	// 2. Читаем тело запроса
	var input struct {
		FilmID int `json:"film_id"`
		Ball   int `json:"ball"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	// 3. Валидация балла (по ТЗ стр. 15: min 1, max 5)
	if input.Ball < 1 || input.Ball > 5 {
		newErrorResponse(w, http.StatusBadRequest, "ball must be between 1 and 5")
		return
	}

	// 4. Вызов сервиса
	id, err := h.services.Rating.Create(userId, input.FilmID, input.Ball)
	if err != nil {
		// Согласно ТЗ (стр. 15), если оценка уже есть, возвращаем 401 и "Score exist"
		if err.Error() == "Score exist" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "invalid",
				"message": "Score exist",
			})
			return
		}
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 5. Успешный ответ (201 Created)
	sendJSON(w, http.StatusCreated, map[string]interface{}{
		"id":     id,
		"status": "success",
	})
}