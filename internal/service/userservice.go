package service

import (
	"errors"
	"gorm.io/gorm"
	"web_application/internal/model"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("field empty")
	}
	return s.db.Create(user).Error
}
