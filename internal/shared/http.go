package shared

import "github.com/gofiber/fiber/v3"

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"name is required"`
}

func ViewData(c fiber.Ctx, data fiber.Map) fiber.Map {
	data["Account"] = c.Locals("account")
	return data
}
