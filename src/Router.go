package main

import (
	"./Middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Use(Middleware.Logger)
	router.Use(Middleware.JsonResponse)

	for _, route := range routes {
		var handler Middleware.ErrorHandlerFunc

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	s := router.PathPrefix("/").Subrouter()
	s.Use(Middleware.Authenticator)

	for _, route := range authenticatedRoutes {
		var handler Middleware.ErrorHandlerFunc
		handler = route.HandlerFunc

		s.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}