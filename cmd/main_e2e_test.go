package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"web_application/internal/handler"
	"web_application/internal/service"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestServer(t *testing.T) *http.ServeMux {
	err := godotenv.Load("/Users/mad/Desktop/dev/go/web_application/.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	service.Migrate(db)

	bookService := service.NewBookService(db)
	bookHandler := handler.NewBookHandler(bookService)

	mux := http.NewServeMux()
	mux.HandleFunc("/books", bookHandler.CreateBook)
	return mux
}

func TestCreateBookE2E(t *testing.T) {
	server := httptest.NewServer(setupTestServer(t))
	defer server.Close()

	body := bytes.NewBufferString(`{"title": "Test Book", "author": "Test Author"}`)
	resp, err := http.Post(server.URL+"/books", "application/json", body)
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestCreateFaliedBookE2E(t *testing.T) {
	server := httptest.NewServer(setupTestServer(t))
	defer server.Close()

	body := bytes.NewBufferString(`{"title": "", "author": "Test Author"}`)
	resp, err := http.Post(server.URL+"/books", "application/json", body)
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
