package usecase

import (
	"errors"
	"testing"

	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/exception"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/repository"
)

var (
	ErrBookSave       = errors.New("book save error")
	ErrAuthorSave     = errors.New("author save error")
	ErrAuthorNotFound = errors.New("author not found")
)

type mockBookRepository struct {
	repository.BookRepository
	saveFunc func(book *entity.Book) error
}

func (m *mockBookRepository) Save(book *entity.Book) error {
	return m.saveFunc(book)
}

type mockAuthorRepository struct {
	repository.AuthorRepository
	saveFunc func(author *entity.Author) error
	findFunc func(name string) (*entity.Author, error)
}

func (m *mockAuthorRepository) FindByName(name string) (*entity.Author, error) {
	return m.findFunc(name)
}

func (m *mockAuthorRepository) Save(author *entity.Author) error {
	return m.saveFunc(author)
}

func TestSuccessCreateBook(t *testing.T) {
	bookRepo := &mockBookRepository{}
	authorRepo := &mockAuthorRepository{}

	setupBookSuccessSave := func(repo *mockBookRepository) {
		repo.saveFunc = func(book *entity.Book) error {
			return nil
		}
	}

	setupAuthorFindSuccess := func(repo *mockAuthorRepository) {
		repo.findFunc = func(name string) (*entity.Author, error) {
			return &entity.Author{Name: "Test Author"}, nil
		}
	}

	setupBookSuccessSave(bookRepo)
	setupAuthorFindSuccess(authorRepo)
	uc := &CreateBook{
		BookRepository:  bookRepo,
		AuthorRepositoy: authorRepo,
		Book:            entity.Book{Title: "Test Book", Author: &entity.Author{Name: "Test Author"}},
	}

	err := uc.Execute()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestSuccessCreateBookWithNewAuthor(t *testing.T) {
	bookRepo := &mockBookRepository{}
	authorRepo := &mockAuthorRepository{}

	setupBookSuccessSave := func(repo *mockBookRepository) {
		repo.saveFunc = func(book *entity.Book) error {
			return nil
		}
	}

	setupAuthorFindNotFound := func(repo *mockAuthorRepository) {
		repo.findFunc = func(name string) (*entity.Author, error) {
			return nil, ErrAuthorNotFound
		}
	}

	setupAuthorSuccessSave := func(repo *mockAuthorRepository) {
		repo.saveFunc = func(author *entity.Author) error {
			return nil
		}
	}

	setupBookSuccessSave(bookRepo)
	setupAuthorFindNotFound(authorRepo)
	setupAuthorSuccessSave(authorRepo)
	uc := &CreateBook{
		BookRepository:  bookRepo,
		AuthorRepositoy: authorRepo,
		Book:            entity.Book{Title: "Test Book", Author: &entity.Author{Name: "Test Author"}},
	}

	err := uc.Execute()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestInvalidData(t *testing.T) {
	uc := &CreateBook{
		Book: entity.Book{Title: "Test Book", Author: &entity.Author{Name: "Test Author"}},
	}

	uc.Book.Title = ""
	err := uc.Execute()
	if err != exception.ErrBookEmptyTitle {
		t.Errorf("expected %v, got %v", exception.ErrBookEmptyTitle, err)
	}
	uc.Book.Title = "Test Book"

	uc.Book.Author = nil
	err = uc.Execute()
	if err != exception.ErrBookEmptyAuthor {
		t.Errorf("expected %v, got %v", exception.ErrBookEmptyAuthor, err)
	}
}

func TestAuthorSaveError(t *testing.T) {
	authorRepo := &mockAuthorRepository{}

	uc := &CreateBook{
		BookRepository:  &mockBookRepository{},
		AuthorRepositoy: authorRepo,
		Book:            entity.Book{Title: "Test Book", Author: &entity.Author{Name: "Test Author"}},
	}

	setupAuthorFindNotFound := func(repo *mockAuthorRepository) {
		repo.findFunc = func(name string) (*entity.Author, error) {
			return nil, ErrAuthorNotFound
		}
	}

	setupAuthorErrorSave := func(repo *mockAuthorRepository) {
		repo.saveFunc = func(author *entity.Author) error {
			return ErrAuthorSave
		}
	}

	setupAuthorFindNotFound(authorRepo)
	setupAuthorErrorSave(authorRepo)
	err := uc.Execute()
	if err != ErrAuthorSave {
		t.Errorf("expected %v, got %v", ErrAuthorSave, err)
	}
}

func TestBookSaveError(t *testing.T) {
	bookRepo := &mockBookRepository{}
	authorRepo := &mockAuthorRepository{}

	uc := &CreateBook{
		BookRepository:  bookRepo,
		AuthorRepositoy: authorRepo,
		Book:            entity.Book{Title: "Test Book", Author: &entity.Author{Name: "Test Author"}},
	}

	setupBookErrorSave := func(repo *mockBookRepository) {
		repo.saveFunc = func(book *entity.Book) error {
			return ErrBookSave
		}
	}
	setupAuthorFindSuccess := func(repo *mockAuthorRepository) {
		repo.findFunc = func(name string) (*entity.Author, error) {
			return &entity.Author{Name: "Test Author"}, nil
		}
	}

	setupAuthorFindSuccess(authorRepo)
	setupBookErrorSave(bookRepo)
	err := uc.Execute()
	if err != ErrBookSave {
		t.Errorf("expected %v, got %v", ErrBookSave, err)
	}
}
