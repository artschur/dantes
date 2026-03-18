package library

import (
	"dantes/internal/books"
	"io"
)

type Library struct {
	Name     string
	Location string
	books    []books.Book
}

func (l *Library) AddBook(book books.Book) {
	l.books = append(l.books, book)
}

func (l *Library) GetBooks() []books.Book {
	return l.books
}

func (l *Library) ReadBook(book books.Book) io.ReadCloser {
	return book.Read()
}
