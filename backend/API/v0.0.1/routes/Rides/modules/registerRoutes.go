package RideRoutes

import (
	"API/routes/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	ridegroup := router.Group("/rides")

	ridegroup.Post("/request", middlewares.JWTVerify, middlewares.RequireRole("client"), middlewares.NoMultipleRide, RideRequest)
	ridegroup.Post("/accept", middlewares.JWTVerify, middlewares.RequireRole("uber"), middlewares.NoMultipleRide, AcceptRide)
	ridegroup.Post("/start", middlewares.JWTVerify, middlewares.RequireRole("uber"), StartRide)
	ridegroup.Post("/finish", middlewares.JWTVerify, middlewares.RequireRole("uber"), FinishRide)

	ridegroup.Post("/getNearRides", middlewares.JWTVerify, middlewares.RequireRole("uber"), GetNearRides)
	ridegroup.Get("/getMyRide", middlewares.JWTVerify, getMyRide)

}
