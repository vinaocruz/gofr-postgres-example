package repository

import (
	"database/sql"

	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/entity"
	"github.com/vinaocruz/gofr-postgres-example/internal/app/library/domain/repository"
	"gofr.dev/pkg/gofr/container"
)

type PostgresBookRepository struct {
	db container.DB
}

func NewPostgresBookRepository(db container.DB) repository.BookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}

func (r *PostgresBookRepository) Save(book *entity.Book) error {
	stmt, err := r.db.Prepare("INSERT INTO books (title, description, author_id, published_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at")
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(book.Title, book.Description, book.Author.ID, book.PublishedAt).Scan(&book.ID, &book.CreatedAt)
}

func (r *PostgresBookRepository) FindAll() (*sql.Rows, error) {
	rows, err := r.db.Query(`SELECT b.id, b.title, b.description, b.published_at, b.created_at, a.id, a.name, a.created_at 
	FROM books b JOIN authors a ON b.author_id = a.id`)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *PostgresBookRepository) Delete(book *entity.Book) error {
	stmt, err := r.db.Prepare("DELETE FROM books WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.ID)
	return err
}
