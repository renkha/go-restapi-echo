package usecase

import (
	"time"

	"github.com/renkha/go-restapi-echo/src/model"
)

type UserResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	AuthToken string     `json:"auth_token"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func UserResponseFormatter(user model.User, authToken string) UserResponse {
	formatter := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		AuthToken: authToken,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return formatter
}
