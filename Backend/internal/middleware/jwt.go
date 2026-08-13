package middleware

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(24 *time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Неверный токен")
	}

	return claims, nil
}

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Токен не передан"})
	}

	tokenStr := strings.TrimPrefix( authHeader, "Bearer ")

	claims, err := ValidateToken(tokenStr)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Неверный токен"})
	}

	c.Locals("user_id", claims["user_id"])

	c.Locals("role", claims["role"])

	return c.Next()
}


func RequireRole (roles ...string) fiber.Handler {
	return func (c *fiber.Ctx) error {
		roleVal := c.Locals("role")
		role, ok := roleVal.(string)

		if !ok || role == "" {
			return c.Status(401).JSON(fiber.Map{"error": "роль не определена"})
		}

		for _, r := range roles {
			if r == role {
				return c.Next()
			} 
		}

		return c.Status(403).JSON(fiber.Map{"error": "для вашей роли доступ запрещен"})
	}


}