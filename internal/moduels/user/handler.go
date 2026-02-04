package user

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}

}
func (h *Handler) GetMe(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "user profile created",
	})
}
