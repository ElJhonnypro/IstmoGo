package UserRoutes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router) {
	authGroup := router.Group("/auth")
	authGroup.Post("/register", RegisterUser)

	//Auth
	if os.Getenv("mode") == "DEV" {
		authGroup.Post("/me", getMe)
	}

}
