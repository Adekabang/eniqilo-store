package routes

import (
	"github.com/Adekabang/eniqilo-store/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, staffController *controller.StaffController) {
	v1Staff := app.Group("/v1/staff")
	v1Staff.Post("/register", staffController.RegisterStaff)
}
