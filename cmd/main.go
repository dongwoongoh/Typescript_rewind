package main

import (
	"net/http"
	"web_application/modules"

	"web_application/internal/service"
)

func main() {
	dsn := modules.ConfigModule()
	db, err := modules.OrmModule(dsn)
	service.Migrate(db)
	bookHandler := modules.BooksModule(db)
	mux := http.NewServeMux()
	mux.HandleFunc("/books", bookHandler.CreateBook)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
