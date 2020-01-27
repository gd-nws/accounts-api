package Middleware

import (
	"../Models"
	"../Services"
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
	"strings"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string {
				"detail": "auth token required",
			})
			return
		}

		splitToken := strings.Split(token, "Bearer")
		token = strings.TrimSpace(splitToken[1])

		claims := &Models.Claims{}
		err := Services.VerifyToken(token, claims)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string {
				"detail": "invalid auth token",
			})
			return
		}

		context.Set(r, "id", claims.Id)
		next.ServeHTTP(w, r)

	})
}