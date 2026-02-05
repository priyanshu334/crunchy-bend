package image

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	type req struct {
		URL       string `json:"url"`
		OwnerType string `json:"owner_type"`
		OwnerID   uint   `json:"owner_id"`
	}
	var body req

	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	if err := h.service.AddImage(body.URL, body.OwnerType, body.OwnerID); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{"success": true})
}

func (h *Handler) List(c *fiber.Ctx) error {
	ownerType := c.Query("owner_type")
	OwnerID := c.QueryInt("owner_id")
	images, err := h.service.GetImages(ownerType, uint(OwnerID))
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.JSON(images)
}
