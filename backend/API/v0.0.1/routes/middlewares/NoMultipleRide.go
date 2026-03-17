package middlewares

import (
	"API/data"
	userUseModels "API/internal/usermanagement/models"

	"github.com/gofiber/fiber/v2"
)

func NoMultipleRide(c *fiber.Ctx) error {
	user := c.Locals("user").(userUseModels.User)
	activeRides, err := data.GetActiveRidesByUserID(data.GetDB(), user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to check active rides",
		})
	}

	if len(activeRides) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "You already have an active ride",
		})
	}
	return c.Next()
}
