package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/reven-erlangga/go-simple-rest-api/controllers"
	"github.com/reven-erlangga/go-simple-rest-api/initializers"
	"github.com/reven-erlangga/go-simple-rest-api/repositories"
	"github.com/reven-erlangga/go-simple-rest-api/routers"
	"github.com/reven-erlangga/go-simple-rest-api/services"
)

func init()  {
	initializers.LoadEnvVariable()
}

func main() {

	initializers.ConnectDatabase()
	initializers.ConnectRedis()

	validate := validator.New()
	bookRepository := repositories.NewBookRepositoryImpl(initializers.DB)
	bookService := services.NewBookServiceImpl(bookRepository, validate)
	bookController := controllers.NewBookController(bookService)
	routes := routers.NewRouter(bookController)
	
	app := fiber.New()
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	// app.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))

	app.Mount("/api", routes)
	
	app.Listen(":3000")
}