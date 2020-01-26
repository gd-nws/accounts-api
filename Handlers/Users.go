package Handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"../Models"
	"../Services"
)

func GetUser(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		return http.StatusUnprocessableEntity, errors.New("id must be an integer")
	}

	userId := context.Get(r, "id").(int)
	if userId != id {
		return http.StatusForbidden, errors.New("cannot access another user's details")
	}

	user, err := Services.GetUserById(id)
	if err != nil {
		return 500, err
	}
	if user.Id == 0 {
		return http.StatusNotFound, errors.New("user not found")
	}

	json.NewEncoder(w).Encode(user)

	return 200, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) (int, error) {
	var newUser Models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusUnprocessableEntity, errors.New("could not process body")
	}

	json.Unmarshal(reqBody, &newUser)

	id, err := Services.CreateUser(newUser)
	if err != nil {
		return 500, err
	}

	newUser.Id = int(id)
	newUser.RefreshToken, err = Services.GenerateToken(newUser, 0)
	if err != nil {
		return 500, errors.New("could not generate token")
	}

	err = Services.UpdateUser(newUser)
	if err != nil {
		return 500, err
	}

	w.Header().Set("Location", "/users/" + strconv.FormatInt(id, 10))
	w.WriteHeader(http.StatusCreated)

	return 201, nil
}