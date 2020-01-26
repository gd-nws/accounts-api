package Handlers

import (
	"../Crypto"
	"../Errors"
	"../Models"
	"../Services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	incorrectCredentials := Errors.NewHttpError(nil, http.StatusUnauthorized, "incorrect credentials")

	if err != nil {
		return Errors.NewHttpError(err, http.StatusUnprocessableEntity, "could not process body")
	}

	var userCredentials Models.UserCredentials
	json.Unmarshal(reqBody, &userCredentials)

	user, err := Services.GetUserByEmail(userCredentials.Email)
	if err != nil {
		clientError, ok := err.(Errors.ClientError) // Check if it is a ClientError.
		if ok {
			status, _ := clientError.ResponseHeaders()
			if status == 404 {
				return incorrectCredentials
			}
		}

		return err
	}

	isMatch := Crypto.ComparePasswords(user.Password, userCredentials.Password)
	if !isMatch {
		return incorrectCredentials
	}

	token, err := Services.GenerateToken(user, time.Hour * 24)
	if err != nil {
		return Errors.NewHttpError(err, http.StatusInternalServerError, "could not generate session token")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {
		"sessionToken": token,
	})

	return nil
}
