package model

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID          uuid.UUID
	Title       string
	Description string
	AuthorName  string
	Year        int
	CreatedAt   time.Time
}

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Author      string `json:"author"`
	Year        int    `json:"year,omitempty"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Author      *string `json:"author,omitempty"`
	Year        *int    `json:"year,omitempty"`
}
