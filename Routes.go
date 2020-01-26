package main

import (
	"./Errors"
	"./Handlers"
	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc rootHandler
}

type rootHandler func(http.ResponseWriter, *http.Request) error
func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)

	if err == nil {
		h := w.Header()
		println(h)
		log.Printf("Status: %d %s", 200, r.RequestURI)
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

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		rootHandler(Handlers.Home),
	},
	Route{
		"GetUser",
		"GET",
		"/users/{id}",
		rootHandler(Handlers.GetUser),
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		rootHandler(Handlers.CreateUser),
	},
	Route{
		"AuthenticateUser",
		"POST",
		"/sessions/login",
		rootHandler(Handlers.Login),
	},
}