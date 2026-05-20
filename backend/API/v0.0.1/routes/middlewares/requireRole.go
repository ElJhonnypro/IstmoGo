package middlewares

import (
	userUseModels "API/internal/usermanagement/models"

	"github.com/gofiber/fiber/v3"
)

func RequireRole(role string) fiber.Handler {
	return func(c fiber.Ctx) error {
		user := c.Locals("user").(userUseModels.User)
		println(user.Role)
		if user.Role != role {
			return c.Status(403).JSON(fiber.Map{
				"error": "Forbidden, you don't have the required role to access this resource, you need: " + role + " and you have: " + user.Role,
			})
		}
		return c.Next()
	}
}
