package routes

import (
	CarModules "API/routes/Car/modules"
	UserModules "API/routes/User/modules"

	"github.com/gofiber/fiber/v2"
)

func StartRoutes() {
	app := fiber.New()

	app.Static("/uploads", "../data/uploads")

	api := app.Group("/api/")

	version := api.Group("/v001")

	// Register User routes
	UserModules.RegisterUserRoutes(version)
	CarModules.RegisterCarRoutes(version)

	// Start the server
	app.Listen(":3000")
}
