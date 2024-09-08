package handler

import (
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

	return uc.Execute()
}
