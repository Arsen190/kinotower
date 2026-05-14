package handler

import (
	"encoding/json"
	"kinotower/internal/domain"
	"kinotower/internal/middleware"
	"net/http"
)

// Регистрация: POST /api/v1/auth/signup
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input domain.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Auth.CreateUser(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusCreated, map[string]interface{}{
		"status": "success",
		"id":     id,
	})
}

// Вход: POST /api/v1/auth/signin
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Auth.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"token":  token,
	})
}

// Удаление аккаунта: DELETE /api/v1/users (ТЗ стр. 12)
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	// Достаем ID пользователя из контекста (положен туда Middleware)
	userId, ok := r.Context().Value(middleware.UserCtx).(int)
	if !ok {
		newErrorResponse(w, http.StatusUnauthorized, "user unauthorized")
		return
	}

	err := h.services.Auth.Delete(userId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Успешное удаление по ТЗ возвращает 204 No Content
	w.WriteHeader(http.StatusNoContent)
}