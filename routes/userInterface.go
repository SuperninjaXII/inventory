package routes

import (
	"suntop/controllers"

	"github.com/gofiber/fiber/v2"
)

func UI(app *fiber.App) {
	app.Get("/productsDashboard", controllers.GetAllProductsDashboard)
	app.Get("/productsPage", controllers.GetAllProductsPages)
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Render("dashboard", fiber.Map{})
	})
	app.Get("/more", func(c *fiber.Ctx) error {
		return c.Render("calculator", fiber.Map{})
	})
	app.Get("/stocklevel", controllers.GetNumberOfProducts)
	app.Get("/schedules", controllers.GetAllSchedules)
}
