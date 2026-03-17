package middlewares

import (
	userUseModels "API/internal/usermanagement/models"

	"github.com/gofiber/fiber/v2"
)

func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(userUseModels.User)

		if user.Role != role {
			return c.Status(403).JSON(fiber.Map{
				"error": "Forbidden",
			})
		}
		return c.Next()
	}
}
