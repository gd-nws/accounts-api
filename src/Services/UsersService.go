package Services

import (
	"../Cache"
	"../Crypto"
	"../Models"
	"../Repositories"
	"errors"
	"github.com/patrickmn/go-cache"
	"strconv"
)

/**
 * Create a user.
 */
func CreateUser(user Models.User) (int, error) {
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
	var user *Models.User
	var err error

	if x, found := Cache.MemoryCache.Get(strconv.Itoa(id)); found {
		user = x.(*Models.User)
	} else {
		if user, err = Repositories.GetUserById(id); err != nil {
			return *user, err
		}
		Cache.MemoryCache.Set(strconv.Itoa(user.Id), user, cache.DefaultExpiration)
	}

	return *user, nil
}

func GetUserByEmail(email string) (Models.User, error) {
	user, err := Repositories.GetUserByEmail(email)

	if err != nil {
		return Models.User{}, err
	}

	return *user, nil
}

func UpdateUser(user Models.User) error {
	err := Repositories.UpdateUser(user)

	if err != nil {
		return err
	}

	return nil
}


func CheckUserCanEdit(userId int, user Models.User) error {
	if userId != user.Id {
		return errors.New("cannot access another user's account")
	}

	return nil
}