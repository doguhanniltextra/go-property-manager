package handler

import (
	"database/sql"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/doguhanniltextra/property_go/internal/service"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	DB *sql.DB
}

func (app *App) CreateRegisterHandler(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given parameteres are invalid",
		})
	}

	getToken, err := service.RegisterService(app.DB, &user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Accepted",
		"status":  201,
		"token":   getToken,
	})
}

func (app *App) CreateAuthHandler(c *fiber.Ctx) error {
	var user model.AuthRequest

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given parameters are invalid",
		})
	}
	b, err := service.AuthService(app.DB, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Given parameters are invalid",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Accepted",
		"bool":    b,
		"name":    user.Name,
	})

}
