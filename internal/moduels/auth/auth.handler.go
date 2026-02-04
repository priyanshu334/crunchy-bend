package auth

import "github.com/gofiber/fiber/v2"

type Hanlder struct {
	service Service
}

func NewHandler(service Service) *Hanlder {
	return &Hanlder{service: service}
}

func (h *Hanlder) Register(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	user, err := h.service.Register(req.Email, req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"id":    user.ID,
		"eamil": user.Email,
	})
}

func (h *Hanlder) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})

	}

	access, refresh, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    access,
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"success": true})
}
