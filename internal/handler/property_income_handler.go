package handler

import (
	"fmt"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/doguhanniltextra/property_go/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (app *App) CreatePropertyIncomeHandler(c *fiber.Ctx) error {
	var propertyIncomeHandler model.PropertyIncome

	if err := c.BodyParser(&propertyIncomeHandler); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given Paramaters Are Invalid",
		})
	}

	result, err := service.PropertyIncomeCreate(app.DB, &propertyIncomeHandler)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given Paramaters Are Invalid",
		})
	}

	row, err := result.RowsAffected()
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given Paramaters Are Invalid",
		})
	}
	fmt.Printf("row: %v\n", row)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Accepted",
		"status":  202,
		"name":    propertyIncomeHandler.PropertyIncomeName,
		"price":   propertyIncomeHandler.PropertyIncomePrice,
	})
}
