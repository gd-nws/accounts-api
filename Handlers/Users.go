package Handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"../Errors"
	"../Models"
	"../Services"
)

func GetUser(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(userId)

	if err != nil {
		return Errors.NewHttpError(nil, http.StatusUnprocessableEntity, "id must be an integer")
	}

	user, err := Services.GetUserById(id)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(user)

	return nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	var newUser Models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Errors.NewHttpError(err, http.StatusUnprocessableEntity, "could not process body")
	}

	json.Unmarshal(reqBody, &newUser)
	id, err := Services.CreateUser(newUser)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", "/users/" + strconv.FormatInt(id, 10))
	w.WriteHeader(http.StatusCreated)

	return nil
}