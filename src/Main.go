package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}