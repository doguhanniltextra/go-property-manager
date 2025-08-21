package middleware

import (
	"fmt"
	"strings"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ReadToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	user, err := ExtractUserFromClaims(claims)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	c.Locals("user", user)
	return c.Next() // artık tip uyumlu
}

// MapClaims içinden güvenli şekilde User objesi çıkarır
func ExtractUserFromClaims(claims jwt.MapClaims) (*model.User, error) {
	// ID
	idFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("id claim missing or invalid")
	}
	id := int(idFloat)

	// Name
	name, ok := claims["name"].(string)
	if !ok {
		return nil, fmt.Errorf("name claim missing or invalid")
	}

	// Email
	email, ok := claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("email claim missing or invalid")
	}

	user := &model.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return user, nil
}
