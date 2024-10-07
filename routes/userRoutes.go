package routes

import (
	"suntop/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	app.Get("/user", controllers.User)
	app.Post("/uploadproduct", controllers.UploadProduct)
	app.Post("/uploadschedule", controllers.UploadSchedule)
	app.Delete("/deleteproduct/:id", controllers.DeleteProductHandler)
	app.Delete("/deleteSchedule/:id", controllers.DeleteScheduleHandler)
	app.Get("/products", controllers.GetAllProducts)
	app.Get("/getproductbyid", controllers.GetProductByID)
}
