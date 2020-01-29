package main

import (
	"./Repositories"
	"log"
	"net/http"
)

func main() {
	Repositories.GetConnection()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}