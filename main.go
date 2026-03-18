package main

import (
	"dantes/internal/books"
	"dantes/internal/library"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	library := library.Library{
		Name: "Arthur's Library",
	}
	library.AddBook(books.Book{Title: "Some Title", Author: "Some author"})
	library.AddBook(books.Book{Title: "SOme title again", Author: "New author"})

	p := tea.NewProgram(library.Model())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
