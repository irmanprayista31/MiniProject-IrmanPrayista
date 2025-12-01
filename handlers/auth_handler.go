package handlers

import (
    "github.com/gofiber/fiber/v2"
    "evermos/models"
    "evermos/services"
)

type AuthHandler struct {
    Service *services.AuthService
}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{Service: services.NewAuthService()}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var payload struct {
        Nama     string `json:"nama"`
        Email    string `json:"email"`
        Password string `json:"password"`
        NoTelp   string `json:"no_telp"`
    }
    if err := c.BodyParser(&payload); err != nil {
        return c.Status(400).JSON(fiber.Map{"error":"invalid payload"})
    }

    user := models.User{
        Nama: payload.Nama,
        Email: payload.Email,
        Password: payload.Password,
        NoTelp: payload.NoTelp,
    }

    var exists models.User
    if err := h.Service.DB.Where("email = ? OR no_telp = ?", user.Email, user.NoTelp).First(&exists).Error; err == nil {
        return c.Status(400).JSON(fiber.Map{"error":"email or phone already used"})
    }

    created, err := h.Service.Register(&user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error":err.Error()})
    }

    created.Password = ""
    return c.Status(201).JSON(created)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var payload struct {
        EmailOrPhone string `json:"email_or_phone"`
        Password     string `json:"password"`
    }
    if err := c.BodyParser(&payload); err != nil {
        return c.Status(400).JSON(fiber.Map{"error":"invalid payload"})
    }
    token, user, err := h.Service.Login(payload.EmailOrPhone, payload.Password)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error":err.Error()})
    }
    user.Password = ""
    return c.JSON(fiber.Map{"token": token, "user": user})
}
