package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-template/internal/patient/controllers"
)

// PublicRoutes func for describe group of public routes.
func PatientRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	route.Post("/patient", controllers.Create)
	route.Get("/patient/:patientMRN", controllers.GetByMRN)
	route.Get("/patient", controllers.GetAll)

}
