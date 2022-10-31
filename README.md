# Book Management System

The book management system allows one to create, retrieve, update and delete books from a database.

The project features the following:

- MySQL Database
- GORM (ORM library)
- Json marshall & unmarshall
- Gorilla Mux (HTTP router and URL matcher)

## Getting Started

### Packages Version

- Go v1.18.1

### Instructions -

#### Local Build

Once you have cloned this repository and installed the go version above:

1. Using the terminal navigate to the root of this project
2. Navigate to the app.go file (`cd pkg/config`) and edit the db configuration (in square brackets) with your username, password & db name (`[user]:[password]@/[dbname]`)
3. Navigate to the main folder folder (`cd cmd/main`)
4. run `go build`
5. run `go run main.go`

#### Local Test

Use POSTMAN or equivalent API to test the following APIs:

- POST "localhost:9000/book" (Create a new book)
  - body: `{"Name": "...", "Author": "...", "Publication": "..."}`
- GET "localhost:9000/books" (Get all books)
- GET "localhost:9000/book/{bookId}" (Get a books by id)
- PUT "localhost:9000/{bookId}" (Update a book)
  - body: `{"Publication": "newPublication"}`
- DELETE "localhost:9000/book/{bookId}" (Delete a book)
