package web

import (
	"demo/internal/shared"

	"github.com/gofiber/fiber/v3"
)

func NotFound(c fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return c.Render("404", shared.ViewData(c, fiber.Map{
		"Title": "Not Found",
	}))
}
