package Services

import (
	"../Crypto"
	"../Errors"
	"../Models"
	"../Repositories"
	"net/http"
)

/**
 * Create a user.
 */
func CreateUser(user Models.User) (int64, error) {
	password, err := Crypto.HashPassword(user.Password)
	if err != nil {
		return 0, Errors.NewHttpError(err, http.StatusInternalServerError, "could not encrypt password")
	}
	user.Password = password

	id, err := Repositories.CreateUser(user)
	if err != nil {
		return 0, Errors.NewHttpError(err, http.StatusInternalServerError, "could not store user")
	}

	return id, err
}

/**
 * Get a user by id.
 */
func GetUserById(id int) (Models.User, error) {
	user, err := Repositories.GetUserById(id)

	if err != nil {
		return Models.User{}, Errors.NewHttpError(err, http.StatusInternalServerError, "could not get user")
	}

	if user.Id == 0 {
		return  Models.User{}, Errors.NewHttpError(nil, http.StatusNotFound, "user not found")
	}

	return user, nil
}

func GetUserByEmail(email string) (Models.User, error) {
	user, err := Repositories.GetUserByEmail(email)
	if err != nil {
		return Models.User{}, Errors.NewHttpError(err, http.StatusInternalServerError, "could not get user")
	}

	if user.Id == 0 {
		return  Models.User{}, Errors.NewHttpError(nil, http.StatusNotFound, "user not found")
	}

	return user, nil
}
