package UserRoutes

import (
	"API/routes/middlewares"
	"os"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router) {
	authGroup := router.Group("/auth")
	authGroup.Post("/register", RegisterUser)
	authGroup.Post("/login", LogIn)

	//Auth
	if os.Getenv("mode") == "DEV" {
		authGroup.Get("/me", middlewares.JWTVerify, getMe)
	}

}
