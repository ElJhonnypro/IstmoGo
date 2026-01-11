package UserRoutes

import (
	"API/internal/userManagement"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	type Request struct {
		Name      string  `json:"name"`
		Phone     string  `json:"phone"`
		Role      string  `json:"role"`
		RID       *string `json:"rid"`
		RIDPhoto  *string `json:"rid_photo"`
		CarID     *string `json:"car_id"`
		Email     string  `json:"email"`
		Password  string  `json:"password"`
		Location  string  `json:"location"`
		Birthdate *string `json:"birthdate"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}
	result := userManagement.RegisterUser(
		req.Name,
		req.Phone,
		req.Role,
		req.RID,
		req.RIDPhoto,
		req.CarID,
		req.Email,
		req.Password,
		req.Location,
		req.Birthdate,
	)

	if result.Success == false {
		return c.Status(fiber.StatusInternalServerError).JSON(result)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
