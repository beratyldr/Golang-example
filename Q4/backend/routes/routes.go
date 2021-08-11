package routes

import (
	"example.com/user/hello/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/products", controllers.GetAll)
	app.Get("/product/:id", controllers.GetProduct)
	app.Post("/create", controllers.New)
	app.Delete("/delete/:id", controllers.Delete)
	app.Put("/edit/:id", controllers.Edit)
}
