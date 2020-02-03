package Handlers

import (
	"../Crypto"
	"../Models"
	"../Services"
	"encoding/json"
	"errors"
	"github.com/gorilla/context"
	"io/ioutil"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) (int, error) {
	var reqBody []byte
	var err error
	incorrectCredentials := errors.New("incorrect credentials")
	unprocessable := errors.New("could not process body")

	if reqBody, err = ioutil.ReadAll(r.Body); err != nil {
		return http.StatusUnprocessableEntity, unprocessable
	}

	var userCredentials Models.UserCredentials
	if err = json.Unmarshal(reqBody, &userCredentials); err != nil {
		return http.StatusUnprocessableEntity, err
	}

	var user Models.User
	if user, err = Services.GetUserByEmail(userCredentials.Email); err != nil {
		return http.StatusInternalServerError, err
	}
	if err = checkIsValidUser(user); err != nil {
		return http.StatusUnauthorized, incorrectCredentials
	}

	if isMatch := Crypto.ComparePasswords(user.Password, userCredentials.Password); !isMatch {
		return  http.StatusUnauthorized, incorrectCredentials
	}

	token, err := Services.GenerateToken(user, time.Hour * 24)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {
		"sessionToken": token,
	})

	return http.StatusOK, nil
}

func RefreshToken(w http.ResponseWriter, r *http.Request) (int, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusUnprocessableEntity, errors.New("could not process body")
	}

	var payload Models.RefreshPayload
	json.Unmarshal(reqBody, &payload)

	claims := &Models.Claims{}
	if err = Services.VerifyToken(payload.RefreshToken, claims); err != nil {
		return http.StatusUnauthorized, err
	}

	token, err := Services.GenerateToken(Models.User{
		Id: claims.Id,
	}, time.Hour * 24)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {
		"sessionToken": token,
	})

	return http.StatusOK, nil
}

func GetSession(w http.ResponseWriter, r *http.Request) (int, error) {
	id := context.Get(r, "id").(int)
	var user Models.User
	var err error

	if user, err = Services.GetUserById(id); err != nil {
		return http.StatusInternalServerError, err
	}
	if err = checkIsValidUser(user); err != nil {
		return http.StatusNotFound, err
	}

	json.NewEncoder(w).Encode(Models.NewUserResponse(user))
	return http.StatusOK, nil
}

func checkIsValidUser(user Models.User) error {
	if user.Id != 0 {
		return nil
	}

	return errors.New("user not found")
}
