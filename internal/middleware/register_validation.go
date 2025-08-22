package middleware

import (
	"net/mail"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/gofiber/fiber/v2"
)

func RegisterValidation(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Given parameters are invalid"})
	}
	if !emailValid(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email is not valid"})
	}
	if !passwordvalid(user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Password type is not valid"})
	}

	return c.Next()
}

func emailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func passwordvalid(password string) bool {
	if password == "" {
		return false
	}
	if len(password) < 4 {
		return false
	}
	return true
}
