package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v4"
	"github.com/lpernett/godotenv"
)

func checkMiddleWare(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf("URL = %s, Method = %s, Time = %s \n", c.OriginalURL(), c.Method(), start)
	return c.Next()
}

// https://youtu.be/nfSbFk8y4D0?si=TIgPOdgUtIcJ91kU&t=4606
// run by nodemon --exec go run main.go route.go --signal SIGTERM
func main() {
	app := fiber.New()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	for i := 0 ; i < 10 ; i++ {
		books = append(books, Book{ ID: len(books) + 1, Title: "Book #" + strconv.Itoa(len(books) + 1), Author: "New" })
	}

	app.Post("/login", logIn)

	// Middleware
	app.Use(checkMiddleWare)

	// jwt
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_ENV"))},
	}))

	// CRUD

	app.Get("/books", getAllBook)

	app.Get("/book/:id", getBook)

	app.Post("/book", createBook)

	app.Put("/book/:id", updateBook)

	app.Delete("/book/:id", deleteBook)

	app.Post("/file", uploadFile)

	app.Get("/env", getEnv)

	app.Listen(":8080")
}
