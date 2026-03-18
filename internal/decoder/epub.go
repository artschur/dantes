package decoder

import (
	"archive/zip"
	"bytes"
	"dantes/internal/books"
	"io"
	"os"

	"github.com/mathieu-keller/epub-parser"
)

type Epub struct {
	defaultSize int64
}

func (e *Epub) Read(book books.Book) (readCloser io.ReadCloser, closer func()error) {
	fileBytes, err := os.ReadFile(book.Path)
	if err != nil {
		return
	}

	bytesReader := bytes.NewReader(fileBytes)
	zipReader, err := zip.NewReader(bytesReader, book.Size)
	if err != nil {
		return
	}

	bookR, err := epub.OpenBook(zipReader)
	if err != nil {
		return
	}
	readCloser, err := bookR.Open(bookR.Container.Rootfile.Path)
	if err != nil {
		return
	}
	return readCloser, bookR.Close
}
