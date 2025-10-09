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
