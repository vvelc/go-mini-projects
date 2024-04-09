package book

import (
	"errors"
	"math/rand"
	"slices"
	"structured-api/pkg/err"
)

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []*Book = []*Book{
	{
		ID:     1,
		Title:  "Book 1",
		Author: "Author 1",
	},
	{
		ID:     2,
		Title:  "Book 2",
		Author: "Author 2",
	},
	{
		ID:     3,
		Title:  "Book 1",
		Author: "Author 3",
	},
}

func AddBook(book *Book) {
	book.ID = rand.Int63n(100)
	books = append(books, book)
}

func GetBookById(id int64) (*Book, error) {
	i := slices.IndexFunc(books, func(b *Book) bool { return b.ID == id })

	if i < 0 {
		return nil, errors.New(err.BookNotFoundFailure)
	}

	return books[i], nil
}

func GetAllBooks() []*Book {
	return books
}

func EditBook(id int64, title string, author string) error {
	affected := 0
	for _, book := range books {
		if book.ID == id {
			book.Title = title
			book.Author = author
			affected = 1
			break
		}
	}

	if affected == 0 {
		return errors.New(err.BookNotFoundFailure)
	}

	return nil
}

func DeleteBook(id int64) error {
	i := slices.IndexFunc(books, func(b *Book) bool { return b.ID == id })

	if i < 0 {
		return errors.New(err.BookNotFoundFailure)
	}

	books = append(books[:i], books[i+1:]...)

	return nil
}
