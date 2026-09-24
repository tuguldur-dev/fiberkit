package auth

import (
	"errors"
	"strings"

	"demo/internal/database"
	"demo/internal/shared"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email" form:"email" example:"admin@example.com"`
	Password string `json:"password" form:"password" example:"admin"`
}

type AccountResponse struct {
	Success bool     `json:"success" example:"true"`
	Account *Account `json:"account"`
}

func LoginPage(c fiber.Ctx) error {
	if CurrentAccount(c) != nil {
		return c.Redirect().To("/")
	}

	return c.Render("auth/login", shared.ViewData(c, fiber.Map{
		"Title": "Login",
		"Error": c.Query("error") != "",
	}))
}

func LoginForm(c fiber.Ctx) error {
	if err := login(c, c.FormValue("email"), c.FormValue("password")); err != nil {
		return c.Redirect().To("/login?error=1")
	}

	return c.Redirect().To("/")
}

func Logout(c fiber.Ctx) error {
	if sess := session.FromContext(c); sess != nil {
		_ = sess.Reset()
	}

	if strings.HasPrefix(c.Path(), "/api/") {
		return c.JSON(fiber.Map{"success": true})
	}

	return c.Redirect().To("/login")
}

// APILogin godoc
// @Summary      Login
// @Description  Start a session cookie (session_id)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginInput  true  "Login credentials"
// @Success      200          {object}  AccountResponse
// @Failure      400          {object}  shared.ErrorResponse
// @Failure      401          {object}  shared.ErrorResponse
// @Router       /login [post]
func APILogin(c fiber.Ctx) error {
	input := new(LoginInput)
	if err := c.Bind().Body(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(shared.ErrorResponse{
			Success: false,
			Error:   "invalid request body",
		})
	}

	if err := login(c, input.Email, input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(shared.ErrorResponse{
			Success: false,
			Error:   "invalid email or password",
		})
	}

	return c.JSON(AccountResponse{
		Success: true,
		Account: CurrentAccount(c),
	})
}

// Me godoc
// @Summary      Current account
// @Description  Get the logged-in account
// @Tags         auth
// @Produce      json
// @Success      200  {object}  AccountResponse
// @Failure      401  {object}  shared.ErrorResponse
// @Security     CookieAuth
// @Router       /me [get]
func Me(c fiber.Ctx) error {
	return c.JSON(AccountResponse{
		Success: true,
		Account: CurrentAccount(c),
	})
}

func login(c fiber.Ctx, email, password string) error {
	email = strings.TrimSpace(email)
	if email == "" || password == "" {
		return errors.New("missing credentials")
	}

	var account Account
	if err := database.DB.Where("email = ?", email).First(&account).Error; err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(password)); err != nil {
		return err
	}

	sess := session.FromContext(c)
	if sess == nil {
		return errors.New("session unavailable")
	}

	if err := sess.Regenerate(); err != nil {
		return err
	}

	sess.Set(sessionAccountID, account.ID)
	c.Locals("account", &account)
	return nil
}
