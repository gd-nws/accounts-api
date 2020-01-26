package Handlers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) (int, error) {
	json.NewEncoder(w).Encode("hello")

	return http.StatusOK, nil
}
