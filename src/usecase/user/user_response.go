package usecase

import (
	"github.com/renkha/go-restapi/src/model"
)

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserResponseFormatter(user model.User) UserResponse {
	formatter := UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return formatter
}
