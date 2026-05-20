package UserRoutes

import (
	userUseModels "API/internal/usermanagement/models"

	"github.com/gofiber/fiber/v3"
)

func getMe(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(c.Locals("user").(userUseModels.User))
}
