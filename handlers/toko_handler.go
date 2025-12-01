package handlers

import (
	"evermos/services"
	"github.com/gofiber/fiber/v2"
)

type TokoHandler struct {
	Service services.TokoService
}

func NewTokoHandler(service services.TokoService) *TokoHandler {
	return &TokoHandler{Service: service}
}

func (h *TokoHandler) Create(c *fiber.Ctx) error {
	var input services.TokoInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	toko, err := h.Service.Create(input)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(toko)
}

func (h *TokoHandler) GetAll(c *fiber.Ctx) error {
	toko, err := h.Service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(toko)
}
