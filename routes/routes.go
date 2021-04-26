package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Home)
}
