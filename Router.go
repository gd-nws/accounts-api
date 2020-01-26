package main

import (
	"./Handlers"
	"./Middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Use(Middleware.Logger)

	for _, route := range routes {
		var handler Middleware.ErrorHandlerFunc

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	s := router.PathPrefix("/users/{id}").Subrouter()
	s.Use(Middleware.Authenticator)
	s.Methods("GET").Path("/").Name("GetUser").Handler(Middleware.ErrorHandlerFunc(Handlers.GetUser))

	return router
}