package books

import (
	"archive/zip"
	"bytes"
	"io"
	"os"

	"github.com/mathieu-keller/epub-parser"
)

func (b *Book) Read() io.ReadCloser {
	fileBytes, err := os.ReadFile(b.Path)
	if err != nil {
		return nil
	}

	bytesReader := bytes.NewReader(fileBytes)
	zipReader, err := zip.NewReader(bytesReader, b.Size)
	if err != nil {
		return nil
	}

	bookR, err := epub.OpenBook(zipReader)
	if err != nil {
		return nil
	}
	readCloser, err := bookR.Open(bookR.Container.Rootfile.Path)
	if err != nil {
		return nil
	}
	return readCloser
}
