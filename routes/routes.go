package routes

import (
    "github.com/gofiber/fiber/v2"

    "evermos/configs"
    "evermos/handlers"
    "evermos/repositories"
    "evermos/services"
)

func Setup(app *fiber.App) {

    api := app.Group("/api")

    authHandler := handlers.NewAuthHandler()
    api.Post("/register", authHandler.Register)
    api.Post("/login", authHandler.Login)

    api.Post("/upload", handlers.UploadFile)

    produkRepo := repositories.NewProdukRepo(configs.DB)
    tokoRepo := repositories.NewTokoRepository(configs.DB)

    produkService := services.NewProdukService(produkRepo)
    tokoService := services.NewTokoService(tokoRepo)

    produkHandler := handlers.NewProdukHandler(produkService)
    tokoHandler := handlers.NewTokoHandler(tokoService)

    api.Post("/produk", produkHandler.Create)
    api.Get("/produk", produkHandler.GetAll)
    api.Get("/produk/:id", produkHandler.GetByID)
    api.Put("/produk/:id", produkHandler.Update)
    api.Delete("/produk/:id", produkHandler.Delete)

    api.Post("/toko", tokoHandler.Create)
    api.Get("/toko", tokoHandler.GetAll)

    app.Static("/uploads", "./uploads")
}
