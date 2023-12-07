package service

import (
	"testing"

	"web_application/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateBookSuccess(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
	bookService := NewBookService(db)
	book := &model.Book{Title: "Test Book", Author: "Test Author"}
	err = bookService.CreateBook(book)
	if err != nil {
		t.Errorf("Failed to create book: %v", err)
	}
}

func TestCreateBookFail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
	bookService := NewBookService(db)
	invalidBook := &model.Book{Title: "", Author: "Test Author"}
	err = bookService.CreateBook(invalidBook)
	if err == nil {
		t.Errorf("Expected an error when creating a book with empty title, got nil")
	} else if err.Error() != "title and author are required" {
		t.Errorf("Expected error message 'title and author are required', got %v", err)
	}
}
