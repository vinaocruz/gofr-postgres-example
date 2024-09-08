package entity

type Book struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PublishedAt string  `json:"published_at"`
	Author      *Author `json:"author"`
	CreatedAt   string  `json:"created_at"`
}
