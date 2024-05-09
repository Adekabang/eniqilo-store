package main

import (
	"github.com/Adekabang/eniqilo-store/controller"
	"github.com/Adekabang/eniqilo-store/db"
	"github.com/Adekabang/eniqilo-store/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	db := db.Connectdb()

	// Controllers
	staffController := controller.NewStaffController(db)

	// Routes
	routes.SetupRoutes(app, staffController)

	app.Listen(":8080")
}
