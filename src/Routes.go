package main

import (
	"./Handlers"
	"./Middleware"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc Middleware.ErrorHandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Middleware.ErrorHandlerFunc(Handlers.Home),
	},

	Route{
		"CreateUser",
		"POST",
		"/users/",
		Middleware.ErrorHandlerFunc(Handlers.CreateUser),
	},
	Route{
		"AuthenticateUser",
		"POST",
		"/sessions/",
		Middleware.ErrorHandlerFunc(Handlers.Login),
	},
	Route{
		"RefreshSession",
		"POST",
		"/sessions/refresh",
		Middleware.ErrorHandlerFunc(Handlers.RefreshToken),
	},
}

var authenticatedRoutes = Routes{
	Route{
		"GetUser",
		"GET",
		"/users/{id}/",
		Middleware.ErrorHandlerFunc(Handlers.GetUser),
	},
	Route{
		"GetSession",
		"GET",
		"/sessions/",
		Middleware.ErrorHandlerFunc(Handlers.GetSession),
	},
}
