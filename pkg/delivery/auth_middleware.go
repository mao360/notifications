package delivery

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type ContextKey string

func (h *Handler) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.sugared.Infof("middleware started")
		header := r.Header.Get("Authorization")
		if header == "" {
			ErrResponseFunc(h.sugared, w, http.StatusUnauthorized, "empty authorization header",
				errors.New("empty authorization header"))
			return
		}
		bisectedHeader := strings.Split(header, " ")
		if bisectedHeader[0] != "Bearer" {
			ErrResponseFunc(h.sugared, w, http.StatusUnauthorized, "not Bearer auth",
				errors.New("not bearer auth"))
			return
		}
		if len(bisectedHeader[1]) == 0 {
			ErrResponseFunc(h.sugared, w, http.StatusUnauthorized, "empty token",
				errors.New("empty token"))
			return
		}
		user, err := h.service.ParseToken(context.Background(), bisectedHeader[1])
		if err != nil {
			ErrResponseFunc(h.sugared, w, http.StatusUnauthorized, "invalid token",
				errors.New("invalid token"))
			return
		}
		h.sugared.Infof("user: %s, %s, %s", user.UserName, user.Password, user.DateOfBirth)
		var key ContextKey
		key = "user"
		ctx := context.WithValue(r.Context(), key, user)
		next.ServeHTTP(w, r.WithContext(ctx))
		h.sugared.Infof("auth middleware completed")
	})
}
