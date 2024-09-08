CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    published_at DATE,
    description TEXT,
    author_id INTEGER,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_author
        FOREIGN KEY(author_id) 
        REFERENCES authors(id)
);