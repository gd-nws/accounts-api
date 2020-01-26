package Handlers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode("hello")

	return nil
}
