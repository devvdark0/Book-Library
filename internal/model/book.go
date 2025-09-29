package model

import "time"

type Book struct {
	ID          int64
	Name        string
	Description string
	Author      string
	Year        uint
	CreatedAt   time.Time
}

type CreateBookRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Year        uint   `json:"year"`
}

type UpdateBookRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Author      *string `json:"author,omitempty"`
	Year        *uint   `json:"year,omitempty"`
}
