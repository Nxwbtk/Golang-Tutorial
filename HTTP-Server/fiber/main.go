package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// https://youtu.be/nfSbFk8y4D0?si=nrHMvbt6BXB8rorp&t=773
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

func getAllBook(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	var book_result Book
	id, err := strconv.Atoi(c.Params("id"))
	if (err != nil) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	for _, book := range books {
		if (book.ID == id) {
			book_result = book
			return c.JSON(book_result)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("ID Not found")
}

func createBook(c *fiber.Ctx) error {
	new_book := new(Book)
	err := c.BodyParser(new_book);
	for _, book := range books {
		if book.ID == new_book.ID {
			return c.Status(fiber.StatusConflict).SendString("Already have")
		}
	}
	if  err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Body not correct")
	}
	books = append(books, *new_book)
	return c.Status(fiber.StatusCreated).JSON(books)
}

func updateBook(c *fiber.Ctx) error {
	updateId, err := strconv.Atoi(c.Params("id"))
	if (err != nil) {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}
	bookBody := new(Book)
	bodyErr := c.BodyParser(bookBody)
	if (bodyErr != nil) {
		return c.Status(fiber.StatusBadRequest).SendString("Body Incorrect")
	}
	for i, book := range books {
		if (book.ID == updateId) {
			books[i].Title = bookBody.Title
			books[i].Author = bookBody.Author
			return c.JSON(books)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not found")
}

func deleteBook(c *fiber.Ctx) error {
	deleteId, err := strconv.Atoi(c.Params("id"))
	if (err != nil) {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}
	for i, book := range books {
		if (book.ID == deleteId) {
			books = append(books[:i], books[i + 1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not found")
}