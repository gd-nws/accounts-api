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
		var email, refreshToken, password string
		var createdAt, updatedAt time.Time
		err = rows.Scan(&id, &email, &password, &refreshToken, &createdAt, &updatedAt)
		if err != nil {
			return []Models.User{}, err
		}

		res = append(res, Models.User{
			Id:           id,
			Password:     password,
			Email:        email,
			RefreshToken: refreshToken,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		})
	}

	return res, err
}

func GetUserById(id int) (Models.User, error) {
	results, err := db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return Models.User{}, err
	}
	defer results.Close()

	users, err := parseUsers(results)

	if len(users) < 1 {
		return Models.User{}, err
	}

	return users[0], err
}

func GetUserByEmail(email string) (Models.User, error) {
	results, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return Models.User{}, err
	}

	defer results.Close()

	users, err := parseUsers(results)

	if len(users) < 1 {
		return Models.User{}, err
	}

	return users[0], err
}

func CreateUser(user Models.User) (int64, error) {
	res, err := db.Exec("INSERT INTO users (email, password, refreshToken) VALUES (?, ?, ?)", user.Email, user.Password, user.RefreshToken)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return id, err
}

func UpdateUser(user Models.User) error {
	_, err := db.Exec(`
		UPDATE users
		SET 
			email = ?,
			refreshToken = ?
		WHERE users.id = ?
	`, user.Email, user.RefreshToken, user.Id)

	return err
}
