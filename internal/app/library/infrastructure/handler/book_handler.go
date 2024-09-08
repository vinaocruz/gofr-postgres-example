package handler

import (
	"strconv"

	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/usecase"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/infrastructure/repository"
	"gofr.dev/pkg/gofr"
)

func CreateBookHandler(ctx *gofr.Context) (interface{}, error) {
	var author entity.Author
	ctx.Bind(&author)

	var book entity.Book
	ctx.Bind(&book)
	book.Author = &author

	uc := usecase.NewCreateBookUseCase(
		repository.NewPostgresBookRepository(ctx.SQL),
		repository.NewPostgresAuthorRepository(ctx.SQL),
		book,
	)

	if err := uc.Execute(); err != nil {
		return nil, err
	}

	return uc.Book, nil
}

func ListBooksHandler(ctx *gofr.Context) (interface{}, error) {
	uc := usecase.NewListBooks(
		repository.NewPostgresBookRepository(ctx.SQL),
	)

	filter := make(map[string]string)
	if name := ctx.Request.Param("name"); name != "" {
		filter["name"] = name
	}

	orderBy := make(map[string]string)
	if order := ctx.Request.Param("order"); order != "" {
		orderBy["published_at"] = order
	}

	return uc.Execute(filter, orderBy)
}

func DeleteBookHandler(ctx *gofr.Context) (interface{}, error) {
	id, _ := strconv.Atoi(ctx.Request.PathParam("id"))

	uc := usecase.NewDeleteBook(
		repository.NewPostgresBookRepository(ctx.SQL),
		id,
	)

	return nil, uc.Execute()
}
