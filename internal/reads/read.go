package reads

import (
	"dantes/internal/books"
	"time"
)

type Read struct {
	Book        books.Book
	LastRead    time.Time
	CurrentPage int
	Completion  int
}

func (r *Read) UpdateCompletion() {
	r.Completion = r.CurrentPage / r.Book.TotalPages
}
