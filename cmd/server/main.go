package main

import (
	"log"
	"os"

	"aygit-muhasebe-integration/config"
	v1 "aygit-muhasebe-integration/internal/api/v1"
	hookv1 "aygit-muhasebe-integration/internal/hook/v1"
	workerv1 "aygit-muhasebe-integration/internal/worker/v1"
	"aygit-muhasebe-integration/pkg/db"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Aygıt Muhasebe Entegrasyonu API
// @version 1.0
// @description Bu servis, NES Özel Entegratör üzerinden e-fatura, e-arşiv ve ürün yönetim süreçlerini yönetir.
// @termsOfService http://swagger.io/terms/

// @contact.name API Destek
// @contact.url http://www.aygit.com/destek
// @contact.email support@aygit.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Initialize configuration
	config.InitConfig()

	// Initialize database connections
	db.InitDB()
	defer db.CloseDB()

	// Seed default data
	db.SeedData()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Aygıt Muhasebe Integration API",
		ErrorHandler: errors.ErrorHandler,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Initialize API V1
	v1.SetupRoutes(app)
	hookv1.SetupHookRoutes(app)

	// Start Workers
	workerv1.StartWorkers()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
