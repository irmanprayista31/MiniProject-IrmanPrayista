package handlers

import (
	"evermos/services"
	"github.com/gofiber/fiber/v2"
)

type ProdukHandler struct {
	Service services.ProdukService
}

func NewProdukHandler(service services.ProdukService) *ProdukHandler {
	return &ProdukHandler{Service: service}
}

func (h *ProdukHandler) Create(c *fiber.Ctx) error {
	var input services.ProdukInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	produk, err := h.Service.Create(input)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(produk)
}

func (h *ProdukHandler) GetAll(c *fiber.Ctx) error {
	produk, err := h.Service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(produk)
}

func (h *ProdukHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	produk, err := h.Service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(produk)
}

func (h *ProdukHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var input services.ProdukInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	produk, err := h.Service.Update(id, input)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(produk)
}

func (h *ProdukHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Service.Delete(id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Produk deleted"})
}
