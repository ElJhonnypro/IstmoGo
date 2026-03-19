package UserRoutes

import (
	carManagement "API/internal/carmanagement"
	userManagement "API/internal/usermanagement"
	"API/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	// Campos normales
	name := c.FormValue("name")
	phone := c.FormValue("phone")
	role := c.FormValue("role")
	email := c.FormValue("email")
	password := c.FormValue("password")
	location := c.FormValue("location")
	birthdate := c.FormValue("birthdate") // string

	// not neccesary
	rid := c.FormValue("rid")
	carID := c.FormValue("carId")

	// Validación básica
	if name == "" || phone == "" || role == "" || email == "" || password == "" || location == "" || birthdate == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Missing required fields",
		})
	}

	// =========================
	// RID PHOTO (file)
	// =========================
	var ridPhotoPath *string = nil
	if role == "uber" {
		// Uber OBLIGATORIO
		if rid == "" || carID == "" {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Missing required uber fields",
			})
		}

		file, err := c.FormFile("photo")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Photo is required",
			})
		}

		photoPath, err := utils.SaveImage(c, file, "rid", "../data/uploads/")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Invalid image",
			})
		}
		ridPhotoPath = &photoPath

		// Validar que el carro exista
		if !carManagement.ValidIdCar(carID) {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Invalid car ID",
			})
		}

	}

	// ❌ Prohibir admin
	if role == "admin" {
		return c.Status(403).JSON(fiber.Map{
			"success": false,
			"message": "Forbidden role",
		})
	}

	if role != "client" && role != "uber" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid role",
		})
	}

	// Convertir opcionales a punteros
	var ridPtr *string = nil
	if rid != "" {
		ridPtr = &rid
	}

	var carIDPtr *string = nil
	if carID != "" {
		carIDPtr = &carID
	}

	var birthdatePtr *string = &birthdate

	// =========================
	// Registrar usuario
	// =========================
	result := userManagement.RegisterUser(
		name,
		phone,
		role,
		ridPtr,
		ridPhotoPath,
		carIDPtr,
		email,
		password,
		location,
		birthdatePtr,
	)

	if !result.Success {
		return c.Status(500).JSON(result)
	}

	return c.Status(200).JSON(result)
}
