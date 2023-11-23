package main

import (
	"github.com/gofiber/fiber/v2"
	messegecontroller "github.com/juanPWT/api-letstalk/controllers/messegeController"
	"github.com/juanPWT/api-letstalk/models"
)

func main() {
	models.ConnectToDb()
	app := fiber.New()

	api := app.Group("/api")

	// route messege
	messege := api.Group("/messege")
	messege.Get("/", messegecontroller.Index)
	messege.Post("/", messegecontroller.Create)

	app.Listen(":8000")
}
