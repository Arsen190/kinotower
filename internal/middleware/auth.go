package middleware

import (
	"context"
	"net/http"
	"strings"
	"kinotower/internal/service"
)

const UserCtx = "userId"

func Auth(services *service.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "empty auth header", http.StatusUnauthorized)
				return
			}

			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				http.Error(w, "invalid auth header", http.StatusUnauthorized)
				return
			}

			if len(headerParts[1]) == 0 {
				http.Error(w, "token is empty", http.StatusUnauthorized)
				return
			}

			userId, err := services.Auth.ParseToken(headerParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserCtx, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}