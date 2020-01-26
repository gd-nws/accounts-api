package Repositories

import (
	"../Models"
	"time"
)

func GetUserById(id int) (Models.User, error) {
	db := getConnection()
	selDB, err := db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return Models.User{}, err
	}

	res := []Models.User{}
	for selDB.Next() {
		var id int
		var email, iv, password string
		var createdAt, updatedAt time.Time
		err = selDB.Scan(&id, &email, &password, &iv, &createdAt, &updatedAt)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, Models.User{
			Id:        id,
			Password:  password,
			Email:     email,
			IV:        iv,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	defer db.Close()
	if len(res) < 1 {
		return Models.User{}, err
	}

	return res[0], err
}

func CreateUser(user Models.User) (int64, error) {
	db := getConnection()
	res, err := db.Exec("INSERT INTO users (email, password, iv) VALUES (?, ?, ?)", user.Email, user.Password, user.IV)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	defer db.Close()
	return id, err
}
