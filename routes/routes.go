package routes

import (
	"github.com/Adekabang/eniqilo-store/controller"
	"github.com/gofiber/fiber/v2"
)

func StaffRoutes(app *fiber.App, staffController *controller.StaffController) {
	v1Staff := app.Group("/v1/staff")
	v1Staff.Post("/register", staffController.RegisterStaff)
	v1Staff.Post("/login", staffController.LoginStaff)
}

func ProductRoutes(app *fiber.App, productController *controller.ProductController) {
	v1Staff := app.Group("/v1/product")
	v1Staff.Post("/", productController.AddProduct)
	v1Staff.Get("/", productController.GetProduct)
	v1Staff.Put("/:id", productController.UpdateProduct)
	v1Staff.Delete("/:id", productController.DeleteProduct)
}
