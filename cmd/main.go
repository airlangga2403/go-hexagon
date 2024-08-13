package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"product-hexagonal-architecture-go/internal/adapters/http"
	"product-hexagonal-architecture-go/internal/application/services"
	"product-hexagonal-architecture-go/pkg/config"
)

func main() {
	// Koneksi ke MongoDB
	client, err := config.ConnectMongoDB()
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	// db name
	db := client.Database("productdb")

	// Setup Fiber
	app := fiber.New()

	// Inisialisasi service dan handler
	productService := services.NewProductService(db)
	productHandler := http.NewProductHandler(productService)

	// Routing
	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products", productHandler.ListProducts)
	app.Get("/products/:id", productHandler.GetProductByID)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)

	// Jalankan server
	log.Fatal(app.Listen(":3000"))
}
