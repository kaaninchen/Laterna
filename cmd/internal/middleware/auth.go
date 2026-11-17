package middleware

import (
	"net/http"

	"github.com/kaaninchen/laterna/server/cmd/internal/db"
	"github.com/kaaninchen/laterna/server/cmd/utils"
)

func WithAdminAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !db.IsAdmin(token) {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		handlerFunc(w, r)
	}
}
