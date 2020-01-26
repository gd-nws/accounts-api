package Repositories

import (
	"../Models"
	"database/sql"
	"time"
)

func parseUsers(rows *sql.Rows) ([]Models.User, error) {
	res := []Models.User{}
	var err error
	for rows.Next() {
		var id int
		var email, iv, password string
		var createdAt, updatedAt time.Time
		err = rows.Scan(&id, &email, &password, &iv, &createdAt, &updatedAt)
		if err != nil {
			return []Models.User{}, err
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

	return res, err
}

func GetUserById(id int) (Models.User, error) {
	db := getConnection()
	results, err := db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		defer db.Close()
		return Models.User{}, err
	}

	users, err := parseUsers(results)

	defer db.Close()
	if len(users) < 1 {
		return Models.User{}, err
	}

	return users[0], err
}

func GetUserByEmail(email string) (Models.User, error) {
	db := getConnection()
	results, err := db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		defer db.Close()
		return Models.User{}, err
	}

	users, err := parseUsers(results)

	defer db.Close()
	if len(users) < 1 {
		return Models.User{}, err
	}

	return users[0], err
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
