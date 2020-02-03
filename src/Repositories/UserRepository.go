package Repositories

import (
	"../Models"
	"database/sql"
	"time"
)

func parseUsers(rows *sql.Rows) ([]Models.User, error) {
	var res []Models.User
	var err error

	for rows.Next() {
		var id int
		var email, refreshToken, password string
		var createdAt, updatedAt time.Time
		if err = rows.Scan(&id, &email, &password, &refreshToken, &createdAt, &updatedAt); err != nil {
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

func GetUserById(id int) (*Models.User, error) {
	const query = `
		SELECT * 
		FROM users 
		WHERE 
			id = $1
	`

	var results *sql.Rows
	var err error
	if results, err = db.Query(query, id); err != nil {
		return &Models.User{}, err
	}
	defer results.Close()

	var users []Models.User
	if users, err = parseUsers(results); err != nil {
		return &Models.User{}, err
	}

	if len(users) < 1 {
		return &Models.User{}, nil
	}

	return &users[0], nil
}

func GetUserByEmail(email string) (*Models.User, error) {
	const query = `
		SELECT * 
		FROM users 
		WHERE 
			email = $1
	`
	var results *sql.Rows
	var err error
	if results, err = db.Query(query, email); err != nil {
		return &Models.User{}, err
	}

	defer results.Close()

	var users []Models.User
	if users, err = parseUsers(results); err != nil {
		return &Models.User{}, err
	}

	if len(users) < 1 {
		return &Models.User{}, nil
	}

	return &users[0], nil
}

func CreateUser(user Models.User) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO users (email, password, refresh_token) 
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.Email, user.Password, user.RefreshToken).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func UpdateUser(user Models.User) error {
	_, err := db.Exec(`
		UPDATE users
		SET 
			email = $1,
			refresh_token = $2
		WHERE users.id = $3
	`, user.Email, user.RefreshToken, user.Id)

	return err
}
