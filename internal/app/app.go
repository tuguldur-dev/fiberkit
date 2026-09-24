package app

import (
	"demo/docs"
	"demo/internal/modules/auth"
	"demo/internal/modules/user"
	"demo/internal/web"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v3"
)

func New() *fiber.App {
	engine := html.New("./views", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	}))
	app.Use(auth.LoadAccount)

	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/api/v1"
	app.Get("/swagger/*", swaggo.HandlerDefault)

	v1 := app.Group("/api/v1")
	auth.Register(app, v1)
	user.Register(app, v1)

	app.Get("/*", static.New("./static/public"))
	app.Use(web.NotFound)

	return app
}
