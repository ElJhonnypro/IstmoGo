package UserRoutes

import (
	"API/internal/userManagement"

	"github.com/gofiber/fiber/v2"
)

func getMe(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	infoUserbyId, err := userManagement.GetUserByID(userId.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error fetching user information",
		})
	}
	return c.Status(fiber.StatusOK).JSON(infoUserbyId)
}
