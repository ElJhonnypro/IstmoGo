package middlewares

import (
	userManagement "API/internal/usermanagement"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func JWTVerify(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid Authorization format",
		})
	}

	tokenStr := strings.TrimSpace(parts[1])

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Token not valid",
		})
	}

	infoUserbyId, err := userManagement.GetUserByID(claims.UserID)
	infoUserbyId.Password = ""

	c.Locals("user", infoUserbyId)

	return c.Next()

}
