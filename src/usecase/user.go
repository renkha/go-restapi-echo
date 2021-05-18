package usecase

import (
	"errors"

	"github.com/renkha/go-restapi/src/model"
	"github.com/renkha/go-restapi/src/repository"
	re "github.com/renkha/go-restapi/src/usecase/user"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(req re.UserRequest) (model.User, error)
	CheckExistEmail(req re.UserRequest) error
	AuthUser(req re.UserLoginRequest) (model.User, error)
}

type usecases struct {
	repository repository.UserRepository
}

func NewUsecase(repository repository.UserRepository) *usecases {
	return &usecases{repository}
}

func (u *usecases) CreateUser(req re.UserRequest) (model.User, error) {
	user := model.User{}
	user.Name = req.Name
	user.Email = req.Email

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
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

func (u *usecases) CheckExistEmail(req re.UserRequest) error {
	email := req.Email

	if user := u.repository.FindEmail(email); user != nil {
		return errors.New("registered email")
	}

	return nil
}

func (u *usecases) AuthUser(req re.UserLoginRequest) (model.User, error) {
	email := req.Email
	password := req.Password

	user, err := u.repository.FindUserByEmail(email)
	if err != nil {
		return user, errors.New("email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid email or password")
	}

	return user, nil
}
