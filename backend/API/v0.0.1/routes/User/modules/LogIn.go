// Process to get Token
package UserRoutes

import (
	userManagement "API/internal/UserManagement"

	"github.com/gofiber/fiber/v2"
)

func LogIn(c *fiber.Ctx) error {
	type Request struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"logIn":   false,
			"message": "Request is wrong.",
		})
	}

	logInResult, token := userManagement.LogIn(req.Identifier, req.Password)

	if logInResult == false {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"logIn":   false,
			"message": "Internal server error.",
		})
	}

	if logInResult {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"logIn":   true,
			"message": "Yay Log in!",
			"Token":   token,
		})

	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"logIn":   false,
		"message": "problem. lol",
	})

}
