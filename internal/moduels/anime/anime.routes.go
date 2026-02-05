package anime

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priyanshu334/go-bend/internal/middleware"
)

func RegisterRoutes(router fiber.Router, hanlder *Handler) {
	anime := router.Group("/anime", middleware.RequireAuth())

	anime.Post("/", hanlder.Create)
	anime.Get("/", hanlder.List)
	anime.Get("/:id", hanlder.Get)
}
