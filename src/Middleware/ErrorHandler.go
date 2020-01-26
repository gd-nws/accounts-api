package Middleware

import (
	"encoding/json"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"time"
)

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) (int, error)
func (fn ErrorHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn(w, r)
	start := context.Get(r, "start").(time.Time)

	if err == nil {
		log.Printf(
			"%d\t%s\t%s\t%s",
			status,
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
		return
	}

	log.Printf(
		"%d\t%s\t%s\t%s\t%s",
		status,
		r.Method,
		r.RequestURI,
		time.Since(start),
		err.Error(),
	)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string {
		"detail": err.Error(),
	})
}