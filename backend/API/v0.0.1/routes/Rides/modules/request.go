package RideRoutes

import (
	rideManagement "API/internal/ridesmanagement"
	userUseModels "API/internal/usermanagement/models"
	"API/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func RideRequest(c *fiber.Ctx) error {
	type Request struct {
		StartLat float64 `json:"start_lat"`
		StartLng float64 `json:"start_lng"`
		EndLat   float64 `json:"end_lat"`
		EndLng   float64 `json:"end_lng"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 🔐 Usuario desde memoria
	user := c.Locals("user").(userUseModels.User)

	// Validar coords básicas
	if req.StartLat == 0 || req.StartLng == 0 || req.EndLat == 0 || req.EndLng == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid coordinates",
		})
	}

	// 📏 Distancia
	distance := utils.Haversine(
		req.StartLat,
		req.StartLng,
		req.EndLat,
		req.EndLng,
	)

	// 💰 Precio
	price := utils.CalculatePrice(distance)

	ride, err := rideManagement.CreateRide(
		user.ID,
		req.StartLat,
		req.StartLng,
		req.EndLat,
		req.EndLng,
		distance,
		price,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":  "Failed to create ride",
			"detail": err.Error(),
		})
	}

	return c.Status(201).JSON(ride)
}
