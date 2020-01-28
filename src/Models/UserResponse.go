package Models

import "time"

type UserResponse struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

/**
 * Get a new instance of a User Response
 */
func NewUserResponse(user User) *UserResponse {
	return &UserResponse{
		Id:           user.Id,
		Email:        user.Email,
		RefreshToken: user.RefreshToken,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
