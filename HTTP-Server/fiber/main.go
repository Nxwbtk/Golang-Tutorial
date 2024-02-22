package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// https://youtu.be/nfSbFk8y4D0?si=nrHMvbt6BXB8rorp&t=773
// run by nodemon --exec go run main.go route.go --signal SIGTERM
func main() {
	app := fiber.New()

	for i := 0 ; i < 10 ; i++ {
		books = append(books, Book{ ID: len(books) + 1, Title: "Book #" + strconv.Itoa(len(books) + 1), Author: "New" })
	}

	// CRUD

	app.Get("/books", getAllBook)

	app.Get("/book/:id", getBook)

	app.Post("/book", createBook)

	app.Put("/book/:id", updateBook)

	app.Delete("/book/:id", deleteBook)

	app.Listen(":8080")
}
