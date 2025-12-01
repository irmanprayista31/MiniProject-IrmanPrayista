package handlers

import (
    "fmt"
    "os"
    "path/filepath"
    "time"

    "github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
    fileHeader, err := c.FormFile("file")
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error":"file required"})
    }

    os.MkdirAll("uploads", os.ModePerm)
    ext := filepath.Ext(fileHeader.Filename)
    filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
    path := filepath.Join("uploads", filename)
    if err := c.SaveFile(fileHeader, path); err != nil {
        return c.Status(500).JSON(fiber.Map{"error":"cannot save file"})
    }
    
    url := "/uploads/" + filename
    return c.JSON(fiber.Map{"url": url})
}
