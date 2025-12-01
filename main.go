package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "evermos/configs"
    "evermos/routes"
    "evermos/models"
)

func main() {
    _ = godotenv.Load()
    configs.InitDB()

    db := configs.DB
    err := db.AutoMigrate(
        &models.User{}, &models.Toko{}, &models.Alamat{}, &models.Category{},
        &models.Produk{}, &models.Foto{}, &models.LogProduk{}, &models.Trx{}, &models.DetailTrx{},
    )
    if err != nil {
        log.Fatal("migrate fail:", err)
    }

    app := fiber.New()

    routes.Setup(app)

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }
    log.Fatal(app.Listen(":" + port))
}
