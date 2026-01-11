package CarModules

import (
	"API/internal/carManagement"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterCarRoutes(router fiber.Router) {
	carGroup := router.Group("/car")
	carGroup.Post("/register", func(c *fiber.Ctx) error {
		// Recibir campos tipo form-data
		plate := c.FormValue("plate")
		model := c.FormValue("model")
		userId := c.FormValue("user_id")
		color := c.FormValue("color")

		// Recibir la imagen (opcional)
		file, err := c.FormFile("photo")
		var filename string = ""

		// Detects if every form value exist
		if plate == "" || model == "" || color == "" || userId == "" {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Missing required fields",
			})
		}

		if err == nil || file != nil {
			filename = uuid.New().String() + "_" + "car" + "_" + file.Filename
			savePath := "../data/uploads/" + filename

			if err := c.SaveFile(file, savePath); err != nil {
				log.Println("Error saving file:", err)
				return c.Status(500).JSON(fiber.Map{
					"success": false,
					"message": "Failed to save image",
				})
			}
			filename = "uploads/" + filename
		} else {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Failed to get image",
			})

		}

		// Crear el auto en DB
		result := carManagement.RegisterCar(
			plate,
			model,
			userId,
			color,
			filename,
		)

		return c.JSON(result)
	})

}
