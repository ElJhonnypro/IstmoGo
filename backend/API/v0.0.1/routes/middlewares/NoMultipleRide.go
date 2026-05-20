package middlewares

import (
	"API/data"
	userUseModels "API/internal/usermanagement/models"

	"github.com/gofiber/fiber/v3"
)

func NoMultipleRide(c fiber.Ctx) error {
	user := c.Locals("user").(userUseModels.User)

	if user.Role == "uber" {
		aRide, err := data.GetRidesByUber(data.GetDB(), user.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to check uber rides",
			})
		}
		if len(aRide) > 0 {
			println("error here tho1")
			return c.Status(400).JSON(fiber.Map{
				"error": "You already have an active ride",
			})
		}
	}

	activeRides, err := data.GetActiveRidesByUserID(data.GetDB(), user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to check active rides",
		})
	}

	if len(activeRides) > 0 {
		println("error here tho2")
		return c.Status(400).JSON(fiber.Map{
			"error": "You already have an active ride",
		})

	}
	return c.Next()
}
