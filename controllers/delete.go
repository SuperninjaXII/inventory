package controllers

import (
	"log"
	"strconv"
	"suntop/config"
	models "suntop/model"

	"github.com/gofiber/fiber/v2"
)

func DeleteProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	productCode, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error converting", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	if err := config.DB.Delete(&models.Product{}, productCode); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": " unable to Delete",
		})
	}
	return c.SendString("deleted")
}

// delete DeleteScheduleHandler
func DeleteScheduleHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	scheduleId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error converting", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	if err := config.DB.Delete(&models.Schedule{}, scheduleId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": " unable to Delete",
		})
	}
	return c.SendString("deleted")
}
