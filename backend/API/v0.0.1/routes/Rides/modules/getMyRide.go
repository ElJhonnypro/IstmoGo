package RideRoutes

import (
	"API/data"
	rideManagement "API/internal/ridesmanagement"
	userModels "API/internal/usermanagement/models"

	"github.com/gofiber/fiber/v2"
)

func getMyRide(c *fiber.Ctx) error {
	user := c.Locals("user").(userModels.User)

	if user.Role == "uber" {
		ride, err := data.GetRidesByUber(data.GetDB(), user.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Get My Rides - Provided by IstmoGo Uber",
			"rides":   ride,
		})
	}

	myRide, err := rideManagement.GetRideByUser(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Get My Ride - Provided by IstmoGo",
		"ride":    myRide,
	})

}
