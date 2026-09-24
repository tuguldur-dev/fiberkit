package user

import (
	"demo/internal/modules/auth"

	"github.com/gofiber/fiber/v3"
)

func Register(app fiber.Router, api fiber.Router) {
	app.Get("/", auth.RequireAuth, Index)
	app.Post("/users", auth.RequireAuth, CreateForm)

	api.Get("/users", auth.RequireAuth, List)
	api.Post("/users", auth.RequireAuth, Create)
}
