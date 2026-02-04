package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/priyanshu334/go-bend/internal/config"
)

func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("access_token")
		if tokenStr == "" {
			return fiber.ErrUnauthorized
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(config.Cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		c.Locals("userID", userID)
		return c.Next()
	}
}

