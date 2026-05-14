package handler

import (
	"kinotower/internal/domain"
	"net/http"
	"strconv"
)

func (h *Handler) getFilms(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	
	page, _ := strconv.Atoi(q.Get("page"))
	if page < 1 { page = 1 }
	
	size, _ := strconv.Atoi(q.Get("size"))
	if size < 1 { size = 10 }
	
	sortBy := q.Get("sortBy")
	if sortBy == "" || sortBy == "rating" { sortBy = "name" } // упрощенно

	sortDir := q.Get("sortDir")
	if sortDir != "desc" { sortDir = "asc" }

	filters := map[string]string{
		"size":    strconv.Itoa(size),
		"offset":  strconv.Itoa((page - 1) * size),
		"sortBy":  sortBy,
		"sortDir": sortDir,
		"search":  q.Get("search"),
		"country": q.Get("country"),
	}

	films, err := h.services.Film.GetAll(filters)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	total, _ := h.services.Film.GetTotalCount(filters)

	sendJSON(w, http.StatusOK, domain.FilmListResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Films: films,
	})
}

func (h *Handler) getFilmById(w http.ResponseWriter, r *http.Request) {
	// Реализация будет позже, пока заглушка
	w.Write([]byte("film detail"))
}