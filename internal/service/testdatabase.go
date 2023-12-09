package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"web_application/internal/model"
)

func SetupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:memdb1?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&model.Book{}, &model.User{}, &model.LoanRecord{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
