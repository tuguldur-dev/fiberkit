package auth

import (
	"strings"

	"fiberkit/internal/database"
	"fiberkit/internal/shared"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

const sessionAccountID = "account_id"

func CurrentAccount(c fiber.Ctx) *Account {
	account, _ := c.Locals("account").(*Account)
	return account
}

func LoadAccount(c fiber.Ctx) error {
	sess := session.FromContext(c)
	if sess == nil {
		return c.Next()
	}

	id := accountID(sess.Get(sessionAccountID))
	if id == 0 {
		return c.Next()
	}

	var account Account
	if err := database.DB.First(&account, id).Error; err == nil {
		c.Locals("account", &account)
	}

	return c.Next()
}

func RequireAuth(c fiber.Ctx) error {
	if CurrentAccount(c) != nil {
		return c.Next()
	}

	if strings.HasPrefix(c.Path(), "/api/") {
		return c.Status(fiber.StatusUnauthorized).JSON(shared.ErrorResponse{
			Success: false,
			Error:   "login required",
		})
	}

	return c.Redirect().To("/login")
}

func accountID(value any) uint {
	switch v := value.(type) {
	case uint:
		return v
	case uint64:
		return uint(v)
	case int:
		if v > 0 {
			return uint(v)
		}
	case int64:
		if v > 0 {
			return uint(v)
		}
	case float64:
		if v > 0 {
			return uint(v)
		}
	}
	return 0
}
