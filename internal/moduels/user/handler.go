package user

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}

}
func (h *Handler) GetMe(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	user, err := h.service.GetByID(userId)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(fiber.Map{
		"id":      user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"message": "user profile created",
	})
}
