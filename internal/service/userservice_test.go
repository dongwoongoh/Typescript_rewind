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

func TestGetUser(t *testing.T) {
	db := SetupDatabase(t)
	userService := NewUserService(db)
	user := &model.User{Name: "Mad123", Email: "integral"}
	err := userService.CreateUser(user)
	if err == nil {
		id := 1
		result, userErr := userService.GetUser(id)
		if userErr == nil {
			t.Logf("user %v:", result)
		}
	}
}

func TestGetUserFail(t *testing.T) {
	db := SetupDatabase(t)
	userService := NewUserService(db)
	_, err := userService.GetUser(1)
	if err.Error() == "error record not found" {
		t.Errorf("error %v", err)
	}
}
