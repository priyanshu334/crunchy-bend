package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, hanlder *Handler) {
	user := router.Group("/auth")
	user.Get("/me", hanlder.GetMe)
}
