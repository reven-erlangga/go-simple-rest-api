package main

import (
	"github.com/reven-erlangga/go-fiber-simple-rest-api/controllers/bookcontroller"
	"github.com/gofiber/fiber/v2"
	"github.com/reven-erlangga/go-fiber-simple-rest-api/models"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	
	book := v1.Group("/books")
	book.Get("/", bookcontroller.index)

	app.Listen(":8000")
}