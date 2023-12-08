package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	messegecontroller "github.com/juanPWT/api-letstalk/controllers/messegeController"
	"github.com/juanPWT/api-letstalk/models"
)

func main() {
	models.ConnectToDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	api := app.Group("/api")

	// route messege
	messege := api.Group("/messege")
	messege.Get("/", messegecontroller.Index)
	messege.Post("/", messegecontroller.Create)

	app.Listen(":8000")
}
