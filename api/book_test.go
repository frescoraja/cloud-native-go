package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Test title", Author: "David Carter", ISBN: "123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Test title","author":"David Carter","isbn":"123456789"}`,
		string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"cloud native go","author":"David Carter","isbn":"091293029"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "cloud native go", Author: "David Carter", ISBN: "091293029"},
		book, "Book JSON unmarshalling wrong.")
}
