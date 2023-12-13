package main

import (
	"fmt"
	"net/http"
	"web_application/modules"

	"web_application/internal/service"
)

type Starter struct {
	message string
	port    int
}

type StarterReturns interface {
	Kit()
}

func (s Starter) Kit() {
	fmt.Printf("cold message: %v, port: %v", s.message, s.port)
	fmt.Println()
}

func main() {
	starter := Starter{message: "web server", port: 8080}
	var starterReturns StarterReturns
	starterReturns = starter
	starterReturns.Kit()
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
