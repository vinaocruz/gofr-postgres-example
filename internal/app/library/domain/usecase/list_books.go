package usecase

import (
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/repository"
)

type ListBooks struct {
	BookRepository repository.BookRepository
}

func NewListBooks(bookRepo repository.BookRepository) *ListBooks {
	return &ListBooks{
		BookRepository: bookRepo,
	}
}

func (uc *ListBooks) Execute() (interface{}, error) {
	var books []entity.Book

	results, err := uc.BookRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var book entity.Book
		var author entity.Author
		if err := results.Scan(
			&book.ID, &book.Title, &book.Description, &book.PublishedAt, &book.CreatedAt,
			&author.ID, &author.Name, &author.CreatedAt); err != nil {
			return nil, err
		}

		book.Author = &author
		books = append(books, book)
	}

	return books, nil
}
