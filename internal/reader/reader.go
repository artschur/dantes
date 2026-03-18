package reader

import (
	"context"
	"dantes/internal/books"
)

type Reader struct {
	context    context.Context
	cancel     context.CancelFunc
	readBuffer []byte
	book       books.Book
}

func NewReader(ctx context.Context, b books.Book, bufferSize int) *Reader {
	ctx, cancel := context.WithCancel(context.Background())
	return &Reader{
		context:    ctx,
		cancel:     cancel,
		readBuffer: make([]byte, bufferSize),
		book:       b,
	}
}
