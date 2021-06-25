package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-template/internal/domain/controllers"
)

// PublicRoutes func for describe group of public routes.
func DomainRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	route.Post("/domain", controllers.Create)
	route.Get("/domain/:id", controllers.GetById)
	route.Get("/domain", controllers.GetAll)

}
