package main

import (
	"./Handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Handlers.Home,
	},
	Route{
		"GetUser",
		"GET",
		"/users/{id}",
		Handlers.GetUser,
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		Handlers.CreateUser,
	},
}