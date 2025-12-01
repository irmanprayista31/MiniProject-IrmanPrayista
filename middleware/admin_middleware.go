package middleware

import (
    "evermos/configs"
    "evermos/models"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
)

func AdminOnly() fiber.Handler {
    return func(c *fiber.Ctx) error {

        userToken := c.Locals("user")
        if userToken == nil {
            return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
        }

        token := userToken.(*jwt.Token)
        claims := token.Claims.(jwt.MapClaims)
        uid := uint(claims["user_id"].(float64))

        var u models.User
        if err := configs.DB.First(&u, uid).Error; err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
        }
        if !u.IsAdmin {
            return c.Status(403).JSON(fiber.Map{"error": "Forbidden, admin only"})
        }

        return c.Next()
    }
}
