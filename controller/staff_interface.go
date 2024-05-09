package controller

import "github.com/gofiber/fiber/v2"

type StaffControllerInterface interface {
	RegisterStaff(*fiber.Ctx) error
}
