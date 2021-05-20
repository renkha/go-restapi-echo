package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/renkha/go-restapi-echo/src/model"
)

type UserRepository interface {
	InsertUser(user model.User) (model.User, error)
	FindEmail(email string) *model.User
	FindUserByEmail(email string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user model.User) (model.User, error) {
	err := r.db.Create(&user)
	if err != nil {
		return user, err.Error
	}

	return user, nil
}

func (r *repository) FindEmail(email string) *model.User {
	var user model.User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil
	}

	return &user
}

func (r *repository) FindUserByEmail(email string) (model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
