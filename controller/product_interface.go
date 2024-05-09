package controller

import "github.com/gofiber/fiber/v2"

type ProductControllerInterface interface {
	AddProduct(*fiber.Ctx) error
	GetProduct(*fiber.Ctx) error
	UpdateProduct(*fiber.Ctx) error
	DeleteProduct(*fiber.Ctx) error
}
