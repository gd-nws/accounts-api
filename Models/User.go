package Models

import "time"

type User struct {
	Id          int       `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IV			string    `json:"iv"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
