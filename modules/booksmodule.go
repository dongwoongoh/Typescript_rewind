package modules

import (
	"gorm.io/gorm"
	"web_application/internal/handler"
	"web_application/internal/service"
)

func BooksModule(db *gorm.DB) *handler.BookHandler {
	bookService := service.NewBookService(db)
	bookHandler := handler.NewBookHandler(bookService)
	return bookHandler
}
