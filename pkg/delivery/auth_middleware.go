package delivery

import (
	"context"
	"net/http"
	"strings"
)

func (h *Handler) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		bisectedHeader := strings.Split(header, " ")
		if bisectedHeader[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if len(bisectedHeader[1]) == 0 {
			//errors.New("empty token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user, err := h.service.ParseToken(context.Background(), bisectedHeader[1])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
