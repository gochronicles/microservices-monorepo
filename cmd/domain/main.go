package main

import (
	"go-fiber-template/internal/domain/routes"
	"os"

	"go-fiber-template/pkg/configs"
	"go-fiber-template/pkg/middleware"
	"go-fiber-template/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "go-fiber-template/docs"// load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.DomainRoutes(app)

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
