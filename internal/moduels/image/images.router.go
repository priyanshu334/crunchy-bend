package image

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priyanshu334/go-bend/internal/middleware"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	img := router.Group("/images", middleware.RequireAuth())
	img.Post("/", handler.Create)
	img.Get("/", handler.List)
}
