package service

import (
	"testing"
	"web_application/internal/model"
)

func TestCreateUser(t *testing.T) {
	db := SetupDatabase(t)
	userService := NewUserService(db)
	user := &model.User{Name: "Mad", Email: "integral"}
	err := userService.CreateUser(user)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
}

func TestCreateUserFail(t *testing.T) {
	db := SetupDatabase(t)
	userService := NewUserService(db)
	invalidUser := &model.User{Name: "MAD", Email: ""}
	err := userService.CreateUser(invalidUser)
	if err.Error() != "field empty" {
		t.Errorf("message %v", err)
	}
}
