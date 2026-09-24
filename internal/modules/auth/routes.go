package auth

import "github.com/gofiber/fiber/v3"

func Register(app fiber.Router, api fiber.Router) {
	app.Get("/login", LoginPage)
	app.Post("/login", LoginForm)
	app.Post("/logout", Logout)

	api.Post("/login", APILogin)
	api.Post("/logout", Logout)
	api.Get("/me", RequireAuth, Me)
}
