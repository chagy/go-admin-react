package controllers

import (
	"strconv"

	"../database"
	"../models"
	"github.com/gofiber/fiber"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
