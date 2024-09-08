package usecase

import (
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/exception"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/repository"
)

type DeleteBook struct {
	BookRepository repository.BookRepository
	id             int
}

func NewDeleteBook(bookRepo repository.BookRepository, id int) *DeleteBook {
	return &DeleteBook{
		BookRepository: bookRepo,
		id:             id,
	}
}

func (uc *DeleteBook) Execute() error {
	row, err := uc.BookRepository.Find(uc.id)
	if err != nil {
		return err
	}

	var book entity.Book
	var author entity.Author
	if err := row.Scan(&book.ID, &book.Title, &book.Description, &book.PublishedAt, &book.CreatedAt, &author.ID, &author.Name, &author.CreatedAt); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return exception.ErrBookNotFound

		}

		return err
	}

	book.Author = &author

	return uc.BookRepository.Delete(&book)
}
