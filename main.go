package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
