package repository

import (
	"database/sql"

	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/repository"
	"gofr.dev/pkg/gofr/container"
)

type PostgresAuthorRepository struct {
	db container.DB
}

func NewPostgresAuthorRepository(db container.DB) repository.AuthorRepository {
	return &PostgresAuthorRepository{
		db: db,
	}
}

func (r *PostgresAuthorRepository) Save(author *entity.Author) error {
	stmt, err := r.db.Prepare("INSERT INTO authors (name) VALUES ($1) RETURNING id, created_at")
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(author.Name).Scan(&author.ID, &author.CreatedAt)
}

func (r *PostgresAuthorRepository) FindAll() (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT id, name FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rows, nil
}

func (r *PostgresAuthorRepository) FindByName(name string) (*entity.Author, error) {
	var author entity.Author

	err := r.db.QueryRow("SELECT id, name, created_at FROM authors WHERE name = $1", name).Scan(&author.ID, &author.Name, &author.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &author, nil
}
