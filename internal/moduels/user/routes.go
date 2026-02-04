package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priyanshu334/go-bend/internal/middleware"
)

func RegisterRoutes(router fiber.Router, hanlder *Handler) {
	user := router.Group("/auth", middleware.RequireAuth())
	user.Get("/me", hanlder.GetMe)
}
