package main

import (
	"fmt"

	"github.com/BILRIK/API/database"
	"github.com/Bilrik/API/book"

	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
}

func helloworld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setuproutes(app *fiber.App) {
	app.Get("/", helloworld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setuproutes(app)

	app.Listen(8080)
}
