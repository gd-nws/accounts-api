package Handlers

import (
	"../Models"
	"../Services"
	"encoding/json"
	"errors"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return http.StatusUnprocessableEntity, errors.New("id must be an integer")
	}

	user, err := Services.GetUserById(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err = Services.CheckUserCanEdit(context.Get(r, "id").(int), user); err != nil {
		return http.StatusForbidden, err
	}

	if err = throwIfNotFound(user); err != nil {
		return http.StatusNotFound, err
	}

	json.NewEncoder(w).Encode(Models.NewUserResponse(user))
	return http.StatusOK, nil
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
		return http.StatusInternalServerError, err
	}

	newUser.Id = id
	newUser.RefreshToken, err = Services.GenerateToken(newUser, 0)
	if err != nil {
		return http.StatusInternalServerError, errors.New("could not generate token")
	}

	err = Services.UpdateUser(newUser)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Location", "/users/" + strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)

	return http.StatusCreated, nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) (int, error) {
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
		return http.StatusInternalServerError, err
	}

	if user.Id == 0 {
		return http.StatusNotFound, errors.New("user not found")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusUnprocessableEntity, errors.New("could not process body")
	}
	json.Unmarshal(reqBody, &user)

	err = Services.UpdateUser(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.WriteHeader(http.StatusNoContent)
	return http.StatusNoContent, nil
}

func throwIfNotFound(user Models.User) error {
	if user.Id != 0 {
		return nil
	}

	return errors.New("user not found")
}