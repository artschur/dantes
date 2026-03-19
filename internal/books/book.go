package books

import (
	"fmt"
	"io"

	"github.com/raitucarp/epub"
	"golang.org/x/net/html"
)

type Book struct {
	Title       string
	Author      string
	Description string
	TotalPages  int
	Location    string
	Path        string
	Size        int64
	DocumentIds []string
	epubReader  *epub.Reader
}

type OutputType int

const (
	OutputMarkdown OutputType = iota
	OutputHTML
)

func (b *Book) Get(w io.Writer, page int, outputType OutputType) error {
	if page < 0 || page >= b.TotalPages {
		return fmt.Errorf("page out of range: %d", page)
	}

	switch outputType {
	case OutputMarkdown:
		node := b.epubReader.ReadContentMarkdownById(b.DocumentIds[page])
		_, err := w.Write([]byte(node))
		return err

	case OutputHTML:
		node := b.epubReader.ReadContentHTMLById(b.DocumentIds[page])
		err := html.Render(w, node)
		return err
	}

	md := b.epubReader.ReadContentHTMLById(b.DocumentIds[page])
	fmt.Println(md)

	return nil
}
