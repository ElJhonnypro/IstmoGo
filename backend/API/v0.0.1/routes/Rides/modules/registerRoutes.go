package RideRoutes

import (
	"API/routes/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	ridegroup := router.Group("/rides")

	ridegroup.Post("/request", middlewares.JWTVerify, RideRequest)

}
