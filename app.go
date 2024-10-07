package main

import (
	"log"
	"suntop/config"
	"suntop/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	config.Database()
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "fiber app",
		})
	})
	routes.User(app)
	routes.UI(app)

	log.Fatal(app.Listen(":3000"))
}
