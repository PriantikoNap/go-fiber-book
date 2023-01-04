package main

import (
	"github.com/PriantikoNap/go-fiber-book.git/controllers"
	"github.com/PriantikoNap/go-fiber-book.git/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBConnection()
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/index", controllers.Index)
	v1.Post("/create", controllers.Create)
	v1.Get("/:id", controllers.Show)
	v1.Put("/:id", controllers.Update)
	v1.Delete("/:id", controllers.Delete)

	app.Listen(":8081")

}
