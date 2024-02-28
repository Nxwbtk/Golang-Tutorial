package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books []Book

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

// use with form-data and key is "image"
func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if (err != nil) {
		return c.Status(fiber.StatusBadRequest).SendString((err.Error()))
	}
	err = c.SaveFile(file, "./uploads/" + file.Filename)
	if (err != nil) {
		return c.Status(fiber.StatusInternalServerError).SendString((err.Error()))
	}
	return c.SendString("File uploaded")
}

type Users struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var user = Users{
	Email: "a@a.com",
	Password: "123",
}

func logIn(c *fiber.Ctx) error {
	usr := new(Users)
	err := c.BodyParser(usr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Body incorrect")
	}
	if user.Email != usr.Email || user.Password != usr.Password {
		return fiber.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = usr.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_ENV")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"token": t,
	})
}