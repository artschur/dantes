package library

import "dantes/internal/books"

type Library struct {
	Name     string
	Location string
	books    []books.Book
}
