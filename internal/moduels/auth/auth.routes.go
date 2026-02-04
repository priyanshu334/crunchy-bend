package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Hanlder) {
	auth := router.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}
