package Handlers

import (
	"encoding/json"
	"github.com/gorilla/context"
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
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		return Errors.NewHttpError(nil, http.StatusUnprocessableEntity, "id must be an integer")
	}

	userId := context.Get(r, "id").(int)

	user, err := Services.GetUserById(userId, id)
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
	if err != nil {
		return Errors.NewHttpError(err, http.StatusInternalServerError, "could not generate token")
	}

	id, err := Services.CreateUser(newUser)
	if err != nil {
		return err
	}

	newUser.Id = int(id)
	newUser.RefreshToken, err = Services.GenerateToken(newUser, 0)
	err = Services.UpdateUser(newUser)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", "/users/" + strconv.FormatInt(id, 10))
	w.WriteHeader(http.StatusCreated)

	return nil
}