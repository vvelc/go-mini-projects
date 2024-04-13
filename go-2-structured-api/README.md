# structured-api

This is the simple-api upgraded, with:
* Standard Go project structure - [go project structure](https://github.com/golang-standards/project-layout)
* Middlewares
* Advanced Logger - [zerolog](https://github.com/rs/zerolog)
* Basic Security features
* Unit Tests - [Testify](https://github.com/stretchr/testify)
* Multiple `Book` entity endpoints

> 💡 `Book` entity endpoints were made with dummy methodologies, without connecting to a database. <br>
> To see how to connect to a database, please go to the next project "crud-api"

## Routes

| Name          | HTTP Method   | Route             |
|---------------|---------------|-------------------|
| Health        | GET	        | /healthz          |
| Create Book	| GET	        | /v1/books         |
| List Books	| GET	        | /v1/books         |
| Get Book		| GET	        | /v1/books/{id}         |
| Update Book	| GET	        | /v1/books/{id}         |
| Delete Book	| GET	        | /v1/books/{id}         |

## Book Entity
Just for reference, here's what the `Book` entity looks like
``` go
type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
```