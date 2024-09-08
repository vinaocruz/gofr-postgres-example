package main

import (
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/infrastructure/handler"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/books", handler.ListBooksHandler)
	app.POST("/books", handler.CreateBookHandler)
	app.DELETE("/books/{id}", handler.DeleteBookHandler)

	app.Run()
}
