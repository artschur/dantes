package books

import (
	"fmt"

	"github.com/raitucarp/epub"
)

func NewBook(path string) (*Book, error) {
	r, err := epub.OpenReader(path)
	if err != nil {
		return nil, fmt.Errorf("error opening epub: %w", err)
	}
	return bookFromReader(&r, path), nil
}

func NewBookFromBytes(data []byte) (*Book, error) {
	r, err := epub.NewReader(data)
	if err != nil {
		return nil, fmt.Errorf("error reading epub: %w", err)
	}
	return bookFromReader(&r, ""), nil
}

func bookFromReader(r *epub.Reader, path string) *Book {
	return &Book{
		Path:        path,
		epubReader:  r,
		Author:      r.Author(),
		Title:       r.Title(),
		Description: r.Description(),
		DocumentIds: r.ListContentDocumentIds(),
		TotalPages:  len(r.ListContentDocumentIds()),
	}
}
