package middleware

import (
    "os"
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v3"
)

func JWTProtected() fiber.Handler {
    secret := os.Getenv("JWT_SECRET")
    return jwtware.New(jwtware.Config{
        SigningKey:   []byte(secret),
        ContextKey:   "user",
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
        },
    })
}