package Middleware

import (
	"../Errors"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"time"
)

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error
func (fn ErrorHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)

	start := context.Get(r, "start").(time.Time)

	if err == nil {
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
		return
	}

	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err) // Log the error.

	clientError, ok := err.(Errors.ClientError) // Check if it is a ClientError.
	if !ok {
		log.Printf("Status: %d %s %v", 500, r.RequestURI, err.Error())
		// If the error is not ClientError, assume that it is ServerError.
		w.WriteHeader(500) // return 500 Internal Server Error.
		return
	}

	body, err := clientError.ResponseBody() // Try to get response body of ClientError.
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	log.Printf("Status: %d %s %v", status, r.RequestURI, string(body)) // Log the error.

	w.WriteHeader(status)
	w.Write(body)
}