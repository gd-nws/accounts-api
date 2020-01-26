package Handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"../Models"
	"../Services"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(userId)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Models.ErrorResponse{
			Message: "Id must be an integer",
			Trace:   "",
		})
		return
	}

	user, err := Services.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Models.ErrorResponse{
			Message: "Could not fetch user",
			Trace:   err.Error(),
		})
		return
	}

	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Models.ErrorResponse{
			Message: "User not found",
			Trace:   "",
		})
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser Models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Models.ErrorResponse{
			Message: "Could not process request",
			Trace:   "",
		})
		return
	}

	json.Unmarshal(reqBody, &newUser)
	id, err := Services.CreateUser(newUser)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Models.ErrorResponse{
			Message: "Could not create user",
			Trace:   err.Error(),
		})
		return
	}

	w.Header().Set("Location", "/users/" + strconv.FormatInt(id, 10))
	w.WriteHeader(http.StatusCreated)
}