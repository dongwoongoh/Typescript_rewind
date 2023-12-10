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

func (s *UserService) GetUser(id int) (*model.User, error) {
	var user model.User
	result := s.db.First(&user, id)
	return &user, result.Error
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.db.Save(user).Error
}
