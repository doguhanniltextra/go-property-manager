package middleware

import "github.com/gofiber/fiber/v2"

func CreateCookie(c *fiber.Ctx) error {

	c.Cookie(&fiber.Cookie{
		Name:  "user_locale",
		Value: "tr_TR",
		Path:  "/",
	})

	return c.Next()
}
