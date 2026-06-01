package middleware

import (
	"strings"

	"crypto-flow/internal/config"

	"github.com/gofiber/fiber/v2"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

func JWT() fiber.Handler {

	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(
				fiber.Map{
					"message": "missing token",
				},
			)
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwtlib.Parse(
			tokenString,
			func(token *jwtlib.Token) (interface{}, error) {
				return []byte(
					config.Get("JWT_SECRET"),
				), nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(401).JSON(
				fiber.Map{
					"message": "invalid token",
				},
			)
		}

		claims := token.Claims.(jwtlib.MapClaims)

		c.Locals(
			"user_id",
			claims["user_id"],
		)

		return c.Next()
	}
}
