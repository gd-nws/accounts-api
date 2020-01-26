package Middleware

import (
	"github.com/gorilla/context"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		context.Set(r, "start", time.Now())
		next.ServeHTTP(w, r)
	})
}