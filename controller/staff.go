package controller

import (
	"database/sql"

	"github.com/Adekabang/eniqilo-store/model"
	"github.com/Adekabang/eniqilo-store/repository"
	"github.com/Adekabang/eniqilo-store/utils"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	errPhoneNumber := utils.ValidatePhoneNumber(payload.PhoneNumber)
	errName := utils.ValidateName(payload.Name)
	errPassword := utils.ValidatePassword(payload.Password)
	if !errPhoneNumber || !errName || !errPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed"})
	}

	repository := repository.NewStaffRepository(DB)
	register := repository.RegisterStaff(*payload)

	if register {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Staff registered successfully"})
	} else {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": register})
	}
}

func (m *StaffController) LoginStaff(c *fiber.Ctx) error {
	DB := m.Db
	var payload *model.LoginStaff

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	repository := repository.NewStaffRepository(DB)
	login := repository.LoginStaff(*payload)

	if login {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Staff login successfully"})
	} else {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": login})
	}
}
