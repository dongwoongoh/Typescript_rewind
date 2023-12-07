package service

import (
	"web_application/internal/model"

	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *model.Book) error {
	return s.db.Create(book).Error
}

func (s *BookService) GetBook(id uint) (*model.Book, error) {
	var book model.Book
	result := s.db.First(&book, id)
	return &book, result.Error
}

func (s *BookService) UpdateBook(book *model.Book) error {
	return s.db.Save(book).Error
}

func (s *BookService) DeleteBook(id uint) error {
	return s.db.Delete(&model.Book{}, id).Error
}
