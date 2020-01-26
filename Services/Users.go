package Services

import (
	"../Crypto"
	"../Models"
	"../Repositories"
)

/**
 * Create a user.
 */
func CreateUser(user Models.User) (int64, error) {
	password, err := Crypto.HashPassword(user.Password)
	user.Password = password

	id, err := Repositories.CreateUser(user)

	return id, err
}

/**
 * Get a user by id.
 */
func GetUser(id int) (Models.User, error) {
	return Repositories.GetUserById(id)
}
