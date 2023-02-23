package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
	"github.com/reven-erlangga/go-simple-rest-api/controllers"
	"github.com/reven-erlangga/go-simple-rest-api/controllers/bookcontroller"
	"github.com/reven-erlangga/go-simple-rest-api/initializers"
)

func init()  {
	initializers.LoadEnvVariable()
}

func main() {
	router := fiber.New()
	app := gin.Default()
	app.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))

	initializers.ConnectDatabase()
	initializers.ConnectRedis()
	
	api := app.Group("/api")
	v1 := api.Group("/v1")
	
	book := v1.Group("/books")
	book.GET("/", bookcontroller.Index)
	book.GET("/:id", bookcontroller.Show)
	book.POST("/", bookcontroller.Create)
	book.PUT("/:id", bookcontroller.Update)
	book.DELETE("/", bookcontroller.Delete)

	user := v1.Group("/users")
	user.POST("/signup", controllers.SignUp)
	user.POST("/login", controllers.Login)

	app.Run()
}