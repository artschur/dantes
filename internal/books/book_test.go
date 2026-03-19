package books_test

import (
	"bytes"
	"dantes/internal/books"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPage(t *testing.T) {
	book, err := books.NewBook("./test_data/the_odyssey.epub")
	require.NoError(t, err)

	var buf bytes.Buffer
	err = book.Get(&buf, 1, books.OutputHTML)

	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
}
