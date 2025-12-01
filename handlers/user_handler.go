package handlers

import (
	"evermos/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var input services.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user, err := h.Service.Register(input)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var input services.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	token, err := h.Service.Login(input)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}
