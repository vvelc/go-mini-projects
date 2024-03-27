package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	assert := assert.New(t)

	expectedBook := &Book{
		ID:     1,
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams",
	}
	newBook := NewBook(
		expectedBook.ID,
		expectedBook.Title,
		expectedBook.Author,
	)

	assert.Equal(expectedBook, newBook, "NewBook() should return a Book with the same properties")
}

func TestListBooks(t *testing.T) {
	assert := assert.New(t)

	expectedBooks := &books
	booksList, err := ListBooks()

	if assert.NoError(err) {
		assert.Equal(expectedBooks, booksList, "ListBooks() should return actual books list")
	}
}

func TestGetBook(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return correct book", func(t *testing.T) {
		expectedBook := &books[0]
		foundBook, err := GetBook(1)

		if assert.NoError(err) {
			assert.Equal(expectedBook, foundBook, "Book() should return correct book")
		}
	})

	t.Run("Should return 'not found' error if book does not exist", func(t *testing.T) {
		foundBook, err := GetBook(1000000)

		assert.Nil(foundBook,  "Book should be nil")

		assert.Error(err, "not found", "Error should be 'not found'")
	})
}
