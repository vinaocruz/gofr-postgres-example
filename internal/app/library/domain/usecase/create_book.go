package usecase

import (
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/exception"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/repository"
)

type CreateBook struct {
	BookRepository  repository.BookRepository
	AuthorRepositoy repository.AuthorRepository
	Book            entity.Book
}

func NewCreateBookUseCase(
	bookRepo repository.BookRepository,
	authorRepo repository.AuthorRepository,
	book entity.Book) *CreateBook {
	return &CreateBook{
		BookRepository:  bookRepo,
		AuthorRepositoy: authorRepo,
		Book:            book,
	}
}

func (uc *CreateBook) validate() error {
	if uc.Book.Title == "" {
		return exception.ErrBookEmptyTitle
	}

	if uc.Book.Author == nil {
		return exception.ErrBookEmptyAuthor
	}

	return nil
}

func (uc *CreateBook) Execute() (err error) {
	err = uc.validate()
	if err != nil {
		return err
	}

	author, err := uc.AuthorRepositoy.FindByName(uc.Book.Author.Name)
	if err != nil {
		author = uc.Book.Author
		err = uc.AuthorRepositoy.Save(author)
		if err != nil {
			return err
		}
	}
	uc.Book.Author = author

	return uc.BookRepository.Save(&uc.Book)
}
