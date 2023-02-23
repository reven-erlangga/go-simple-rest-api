package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reven-erlangga/go-simple-rest-api/controllers"
)

func NewRouter(c *controllers.BookController) *fiber.App {
	router := fiber.New()
	v1 := router.Group("/v1")

	book := v1.Group("/books")
	book.Get("/", c.FindAll)
	book.Get("/:id", c.FindById)
	book.Post("/", c.Create)
	book.Patch("/:id", c.Update)
	book.Delete("/", c.Delete)

	return router
}

// api := app.Group("/api")
// v1 := api.Group("/v1")

// book := v1.Group("/books")
// book.GET("/", bookcontroller.Index)
// book.GET("/:id", bookcontroller.Show)
// book.POST("/", bookcontroller.Create)
// book.PUT("/:id", bookcontroller.Update)
// book.DELETE("/", bookcontroller.Delete)