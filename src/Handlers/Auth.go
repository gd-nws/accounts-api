package Handlers

import (
	"../Crypto"
	"../Models"
	"../Services"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) (int, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	incorrectCredentials := errors.New("incorrect credentials")

	if err != nil {
		return http.StatusUnprocessableEntity, errors.New("could not process body")
	}

	var userCredentials Models.UserCredentials
	json.Unmarshal(reqBody, &userCredentials)

	user, err := Services.GetUserByEmail(userCredentials.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if user.Id == 0 {
		return http.StatusUnauthorized, incorrectCredentials
	}

	isMatch := Crypto.ComparePasswords(user.Password, userCredentials.Password)
	if !isMatch {
		return  http.StatusUnauthorized, incorrectCredentials
	}

	token, err := Services.GenerateToken(user, time.Hour * 24)
	if err != nil {
		return 500, err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {
		"sessionToken": token,
	})

	return 200, nil
}

func RefreshToken(w http.ResponseWriter, r *http.Request) (int, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusUnprocessableEntity, errors.New("could not process body")
	}

	var payload Models.RefreshPayload
	json.Unmarshal(reqBody, &payload)

	claims := &Models.Claims{}
	err = Services.VerifyToken(payload.RefreshToken, claims)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	token, err := Services.GenerateToken(Models.User{
		Id: claims.Id,
	}, time.Hour * 24)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {
		"sessionToken": token,
	})

	return 200, nil
}
