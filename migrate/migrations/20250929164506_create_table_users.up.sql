CREATE TABLE book (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    author VARCHAR(255) NOT NULL,
    year INTEGER NOT NULL,
    created_at TIMESTAMP
);