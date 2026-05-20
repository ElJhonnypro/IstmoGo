package RideRoutes

import (
	rideManagement "API/internal/ridesmanagement"

	Fiber "github.com/gofiber/fiber/v3"
)

func GetRides(c Fiber.Ctx) error {

	AllRides, err := rideManagement.GetAllRides() // Aquí se podrían obtener los rides desde la base de datos
	if err != nil {
		return c.Status(500).JSON(Fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(Fiber.Map{
		"message": "Get Rides - Provided by IstmoGo",
		"rides":   AllRides, // Aquí se podrían listar los rides disponibles
	})
}

func GetNearRides(c Fiber.Ctx) error {
	type Request struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	var req Request
	if err := c.Bind().Body(&req); err != nil {
		println(err.Error())
		return c.Status(400).JSON(Fiber.Map{
			"error": "Invalid request body",
		})
	}

	lat := req.Lat
	lng := req.Lng

	rides, err := rideManagement.GetRidesNearLocation(lat, lng)
	if err != nil {
		return c.Status(500).JSON(Fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(Fiber.Map{
		"message": "Get Near Rides - Provided by IstmoGo",
		"rides":   rides,
	})
}
