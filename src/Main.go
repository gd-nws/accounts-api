package main

import (
	"./Cache"
	"./Repositories"
	"log"
	"net/http"
)

func main() {
	Repositories.GetConnection()
	Cache.CreateCache()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}