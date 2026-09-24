package user

import (
	"strings"

	"fiberkit/internal/database"
	"fiberkit/internal/shared"

	"github.com/gofiber/fiber/v3"
)

type CreateUserInput struct {
	Name string `json:"name" form:"name" example:"Ada Lovelace"`
}

type UsersResponse struct {
	Success bool   `json:"success" example:"true"`
	Users   []User `json:"users"`
}

type UserResponse struct {
	Success bool `json:"success" example:"true"`
	User    User `json:"user"`
}

func Index(c fiber.Ctx) error {
	var users []User
	if err := database.DB.Order("id desc").Find(&users).Error; err != nil {
		return err
	}

	return c.Render("user/index", shared.ViewData(c, fiber.Map{
		"Title": "Users",
		"Users": users,
	}))
}

func CreateForm(c fiber.Ctx) error {
	name := strings.TrimSpace(c.FormValue("name"))
	if name != "" {
		database.DB.Create(&User{Name: name})
	}

	return c.Redirect().To("/")
}

// List godoc
// @Summary      List users
// @Description  Get all users from SQLite
// @Tags         users
// @Produce      json
// @Success      200  {object}  UsersResponse
// @Failure      401  {object}  shared.ErrorResponse
// @Failure      500  {object}  shared.ErrorResponse
// @Security     CookieAuth
// @Router       /users [get]
func List(c fiber.Ctx) error {
	var users []User
	if err := database.DB.Order("id desc").Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(shared.ErrorResponse{
			Success: false,
			Error:   "could not list users",
		})
	}

	return c.JSON(UsersResponse{
		Success: true,
		Users:   users,
	})
}

// Create godoc
// @Summary      Create user
// @Description  Create a user in SQLite
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      CreateUserInput  true  "User to create"
// @Success      201   {object}  UserResponse
// @Failure      400   {object}  shared.ErrorResponse
// @Failure      401   {object}  shared.ErrorResponse
// @Failure      500   {object}  shared.ErrorResponse
// @Security     CookieAuth
// @Router       /users [post]
func Create(c fiber.Ctx) error {
	input := new(CreateUserInput)
	if err := c.Bind().Body(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(shared.ErrorResponse{
			Success: false,
			Error:   "invalid request body",
		})
	}

	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(shared.ErrorResponse{
			Success: false,
			Error:   "name is required",
		})
	}

	item := User{Name: input.Name}
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(shared.ErrorResponse{
			Success: false,
			Error:   "could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(UserResponse{
		Success: true,
		User:    item,
	})
}
