package controller

import (
	"database/sql"
	"log"

	"github.com/Adekabang/eniqilo-store/model"
	"github.com/Adekabang/eniqilo-store/repository"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	Db *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{Db: db}
}

func (m *ProductController) AddProduct(c *fiber.Ctx) error {

	DB := m.Db
	var payload *model.AddUpdateProduct

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	repository := repository.NewProductRepository(DB)
	addProduct := repository.AddProduct(*payload)

	if addProduct {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Add Product successfully"})
	} else {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": addProduct})
	}
}

func (m *ProductController) GetProduct(c *fiber.Ctx) error {

	reqQuery := c.Queries()
	log.Println("reqQuery", reqQuery)

	DB := m.Db

	// if err := c.BodyParser(&payload); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	// }

	repository := repository.NewProductRepository(DB)
	getProduct := repository.GetProduct(model.ParamsGetProduct{})

	log.Println("uuy")

	if getProduct {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Get Product successfully"})
	} else {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": getProduct})
	}
}

func (m *ProductController) UpdateProduct(c *fiber.Ctx) error {
	DB := m.Db
	id := c.Params("id")
	log.Println("id", id)
	var payload *model.AddUpdateProduct

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	repository := repository.NewProductRepository(DB)
	updateProduct := repository.UpdateProduct(*payload)

	if updateProduct {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Update Product successfully"})
	} else {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": updateProduct})
	}
}

func (m *ProductController) DeleteProduct(c *fiber.Ctx) error {
	DB := m.Db
	id := c.Params("id")
	log.Println("id", id)

	repository := repository.NewProductRepository(DB)
	deleteProduct := repository.DeleteProduct(model.ProductUri{Id: id})

	if deleteProduct {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Delete Product successfully"})
	} else {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": deleteProduct})
	}
}
