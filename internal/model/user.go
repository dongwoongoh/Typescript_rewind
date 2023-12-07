package model

type User struct {
	BaseModel
	Name  string
	Email string
	Books []Book `gorm:"many2many:user_books;"`
}
