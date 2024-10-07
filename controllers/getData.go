package controllers

import (
	"log"
	"strconv"
	"suntop/config"
	"suntop/model"

	"github.com/gofiber/fiber/v2"
)

// Get all products
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		log.Println("error fetching product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error fetching books"})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}

// Get all products dashboard
func GetAllProductsDashboard(c *fiber.Ctx) error {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		log.Println("error fetching product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error fetching books"})
	}
	return c.Status(fiber.StatusOK).Render("table", fiber.Map{
		"product": products,
	})
}

func GetAllProductsPages(c *fiber.Ctx) error {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		log.Println("error fetching product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error fetching books"})
	}
	return c.Status(fiber.StatusOK).Render("product", fiber.Map{
		"product": products,
	})
}

// Search product by title or description

func SearchProduct(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "query parameter is required"})
	}

	var products []models.Product
	err := config.DB.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	if err != nil {
		log.Println("error searching books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}

// get product by id
func GetProductByID(c *fiber.Ctx) error {
	db := config.DB
	id := c.Params("id")
	var product models.Product
	result := db.First(&product, id)
	if result.Error != nil {
		// Handle the error, log it, and return a meaningful response
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(product)
}
func GetNumberOfProducts(c *fiber.Ctx) error {
	var count int64
	if err := config.DB.Model(&models.Product{}).Count(&count).Error; err != nil {
		log.Fatal(err)
	}
	return c.SendString(strconv.FormatInt(count, 10))
}
func GetAllSchedules(c *fiber.Ctx) error {
	var schedule []models.Schedule
	if err := config.DB.Find(&schedule).Error; err != nil {
		log.Println("error fetching schedule:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error fetching books"})
	}
	return c.Status(fiber.StatusOK).Render("schedules", fiber.Map{
		"schedule": schedule,
	})
}
