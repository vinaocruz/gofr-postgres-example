package repository

import (
	"database/sql"

	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
)

type BookRepository interface {
	Save(book *entity.Book) error
	FindAll() (*sql.Rows, error)
	Find(id int) (*sql.Row, error)
	Delete(book *entity.Book) error
}
