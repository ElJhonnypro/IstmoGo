package CarModules

import (
	carManagement "API/internal/carmanagement"
	"API/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterCarRoutes(router fiber.Router) {
	carGroup := router.Group("/car")
	carGroup.Post("/register", func(c *fiber.Ctx) error {
		// Recibir campos tipo form-data
		plate := c.FormValue("plate")
		model := c.FormValue("model")
		color := c.FormValue("color")

		// Recibir la imagen (opcional)

		// Detects if every form value exist
		if plate == "" || model == "" || color == "" {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Missing required fields",
			})
		}

		file, err := c.FormFile("photo")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Photo is required",
			})
		}

		photoPath, err := utils.SaveImage(c, file, "car", "../data/uploads/")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Invalid image",
			})
		}

		result := carManagement.RegisterCar(
			plate,
			model,
			color,
			photoPath,
		)

		// Crear el auto en DB

		return c.JSON(result)
	})

}
