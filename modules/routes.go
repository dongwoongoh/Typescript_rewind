package modules

import (
	"gorm.io/gorm"
	"net/http"
)

func MuxRoutes(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()
	bookHandler := BooksModule(db)
	mux.HandleFunc("/books", bookHandler.CreateBook)
	return mux
}
