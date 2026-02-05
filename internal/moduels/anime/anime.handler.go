package anime

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	type req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	var body req

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	anime, err := h.service.Create(body.Title, body.Description)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(anime)

}

func (h *Handler) List(c *fiber.Ctx) error {
	list, err := h.service.List()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(list)
}

func (h *Handler) Get(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	anime, err := h.service.Get(uint(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(anime)
}
