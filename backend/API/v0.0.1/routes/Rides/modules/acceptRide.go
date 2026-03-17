package RideRoutes

import (
	rideManagement "API/internal/ridesmanagement"

	"github.com/gofiber/fiber/v2"
)

func AcceptRide(c *fiber.Ctx) error {
	type Request struct {
		RideID   string `json:"ride_id"`
		DriverID string `json:"driver_id"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	acceptedRide, err := rideManagement.AcceptRide(req.RideID, req.DriverID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Ride accepted successfully",
		"ride":    acceptedRide,
	})
}
