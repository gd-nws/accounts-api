package Middleware

import (
	"../Models"
	"../Services"
	"github.com/gorilla/context"
	"net/http"
	"strings"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer")
		token = strings.TrimSpace(splitToken[1])

		claims := &Models.Claims{}
		err := Services.VerifyToken(token, claims)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		context.Set(r, "id", claims.Id)
		next.ServeHTTP(w, r)

	})
}