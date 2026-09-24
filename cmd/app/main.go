package main

import (
	"fiberkit/internal/app"
	"fiberkit/internal/database"
	"fiberkit/internal/modules/auth"
	"fiberkit/internal/modules/user"
	"flag"
	"log"

	"github.com/gofiber/fiber/v3"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

// @title           Fiber Template API
// @version         1.0
// @description     Simple Fiber template with HTML includes, Swagger, GORM, SQLite, and session auth
// @schemes         http
// @BasePath        /api/v1
// @securityDefinitions.apikey CookieAuth
// @in cookie
// @name session_id
func main() {
	flag.Parse()
	database.Connect(&auth.Account{}, &user.User{})
	auth.Seed()

	log.Fatal(app.New().Listen(*port, fiber.ListenConfig{EnablePrefork: *prod}))
}
