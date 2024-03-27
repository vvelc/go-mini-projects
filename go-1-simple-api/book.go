package main

import (
	"errors"
	"slices"
)

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
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
		Title:  "Book 3",
		Author: "Author 3",
	},
	{
		ID:     4,
		Title:  "Book 4",
		Author: "Author 4",
	},
	{
		ID:     5,
		Title:  "Book 5",
		Author: "Author 5",
	},
}

func NewBook(id int64, title string, author string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
}

func ListBooks() (*[]Book, error) {
	return &books, nil
}

func GetBook(id int64) (*Book, error) {
	i := slices.IndexFunc(books, func(b Book) bool { return b.ID == id })

	if i < 0 {
		return nil, errors.New("not found")
	}

	return &books[i], nil
}
