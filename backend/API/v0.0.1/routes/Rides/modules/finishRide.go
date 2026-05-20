package RideRoutes

import (
	rideManagement "API/internal/ridesmanagement"

	"github.com/gofiber/fiber/v2"
)

func FinishRide(c *fiber.Ctx) error {
	type Request struct {
		RideID string `json:"ride_id"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body, " + req.RideID,
		})
	}

	finishedRide, err := rideManagement.FinishRide(req.RideID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Ride finished successfully",
		"ride":    finishedRide,
	})
}
