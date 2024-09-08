package exception

import "errors"

var (
	ErrBookNotFound    = errors.New("book not found")
	ErrBookEmptyTitle  = errors.New("book title is empty")
	ErrBookEmptyAuthor = errors.New("book author is empty")
)
