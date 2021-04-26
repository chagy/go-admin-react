package controllers

import "github.com/gofiber/fiber"

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
