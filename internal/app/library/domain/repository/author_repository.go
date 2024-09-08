package repository

import (
	"database/sql"

	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
)

type AuthorRepository interface {
	Save(author *entity.Author) error
	FindAll() (*sql.Rows, error)
	FindByName(name string) (*entity.Author, error)
}
