package model

type Book struct {
}

type CreateBookRequest struct {
	Name        string
	Description string
	Author      string
	CreatedAt   string
}
