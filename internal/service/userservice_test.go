package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web_application/internal/model"
)

func TestCreateUser(t *testing.T) {
	db := SetupDatabase(t)
	userService := NewUserService(db)
	user := &model.User{Name: "Mad", Email: "integral"}
	err := userService.CreateUser(user)
	assert.Nil(t, err)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
}

func TestCreateUserFail(t *testing.T) {
	assert := assert.New(t)
	db := SetupDatabase(t)
	userService := NewUserService(db)
	invalidUser := &model.User{Name: "MAD", Email: ""}
	err := userService.CreateUser(invalidUser)
	assert.NotNil(err)
	assert.EqualError(err, "field empty")
}
