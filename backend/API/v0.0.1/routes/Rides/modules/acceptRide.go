package RideRoutes

import (
	rideManagement "API/internal/ridesmanagement"
	userModels "API/internal/usermanagement/models"

	"github.com/gofiber/fiber/v2"
)

func AcceptRide(c *fiber.Ctx) error {

	User := c.Locals("user").(userModels.User)

	type Request struct {
		RideID string `json:"ride_id"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	acceptedRide, err := rideManagement.AcceptRide(req.RideID, User.ID)
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
