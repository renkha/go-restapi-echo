package usecase

import (
	"errors"

	"github.com/renkha/go-restapi/src/model"
	"github.com/renkha/go-restapi/src/repository"
	usecase "github.com/renkha/go-restapi/src/usecase/user"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(req usecase.UserRequest) (model.User, error)
	CheckExistEmail(req usecase.UserRequest) error
}

type usecases struct {
	repository repository.UserRepository
}

func NewUsecase(repository repository.UserRepository) *usecases {
	return &usecases{repository}
}

func (u *usecases) CreateUser(req usecase.UserRequest) (model.User, error) {
	user := model.User{}
	user.Name = req.Name
	user.Email = req.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)

	newUser, err := u.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (u *usecases) CheckExistEmail(req usecase.UserRequest) error {
	email := req.Email

	if user := u.repository.FindEmail(email); user != nil {
		return errors.New("registered email")
	}

	return nil
}
