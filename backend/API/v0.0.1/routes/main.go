package routes

import (
	CarModules "API/routes/Car/modules"
	RideRoutes "API/routes/Rides/modules"
	UserModules "API/routes/User/modules"

	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartRoutes() {

	app := fiber.New()

	// =========================
	// Create log file
	// =========================
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	// =========================
	// Logger Middleware
	// =========================
	app.Use(func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		log := fmt.Sprintf(
			"%s | %s | %s | %d | %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Method(),
			c.OriginalURL(),
			c.Response().StatusCode(),
			duration,
		)

		// Terminal log
		fmt.Print(log)

		// File log
		logFile.WriteString(log)

		return err
	})

	// =========================
	// CORS
	// =========================
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// =========================
	// Static files
	// =========================
	app.Static("/uploads", "../data/uploads")

	api := app.Group("/api/")
	version := api.Group("/v001")

	// =========================
	// Routes
	// =========================
	UserModules.RegisterUserRoutes(version)
	CarModules.RegisterCarRoutes(version)
	RideRoutes.RegisterRoutes(version)

	// =========================
	// Start server
	// =========================
	app.Listen(":2534")
}
