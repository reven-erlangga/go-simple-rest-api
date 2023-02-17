package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reven-erlangga/go-fiber-simple-rest-api/controllers/bookcontroller"
	"github.com/reven-erlangga/go-fiber-simple-rest-api/models"
)

func main() {
	app := gin.Default()
	models.ConnectDatabase()
	
	api := app.Group("/api")
	v1 := api.Group("/v1")
	
	book := v1.Group("/books")
	book.GET("/", bookcontroller.Index)
	book.GET("/:id", bookcontroller.Show)
	book.POST("/", bookcontroller.Create)
	book.PUT("/:id", bookcontroller.Update)
	book.DELETE("/:id", bookcontroller.Delete)

	app.Run()
}