package Services

import (
	"../Crypto"
	"../Models"
	"../Repositories"
	"errors"
)

/**
 * Create a user.
 */
func CreateUser(user Models.User) (int64, error) {
	password, err := Crypto.HashPassword(user.Password)
	if err != nil {
		return 0, errors.New("could not encrypt password")
	}
	user.Password = password

	id, err := Repositories.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return id, err
}

/**
 * Get a user by id.
 */
func GetUserById(id int) (Models.User, error) {

	user, err := Repositories.GetUserById(id)

	if err != nil {
		return Models.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (Models.User, error) {
	user, err := Repositories.GetUserByEmail(email)

	if err != nil {
		return Models.User{}, err
	}

	return user, nil
}

func UpdateUser(user Models.User) error {
	err := Repositories.UpdateUser(user)

	if err != nil {
		return err
	}

	return nil
}
