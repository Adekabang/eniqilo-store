package controller

import (
	"database/sql"

	"github.com/Adekabang/eniqilo-store/model"
	"github.com/Adekabang/eniqilo-store/repository"
	"github.com/gofiber/fiber/v2"
)

type StaffController struct {
	Db *sql.DB
}

func NewStaffController(db *sql.DB) *StaffController {
	return &StaffController{Db: db}
}

func (m *StaffController) RegisterStaff(c *fiber.Ctx) error {
	DB := m.Db
	var payload *model.RegisterStaff

	if err := c.BodyParser(&payload); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	repository := repository.NewStaffRepository(DB)
	register := repository.RegisterStaff(*payload)
	if register {
		// c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": newNote}})

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
	} else {

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": register})
	}
}
