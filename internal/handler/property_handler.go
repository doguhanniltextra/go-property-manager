package handler

import (
	"fmt"
	"strconv"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/doguhanniltextra/property_go/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (app *App) CreatePropertyHandler(c *fiber.Ctx) error {
	var property model.Property

	userVal := c.Locals("user")
	user, ok := userVal.(*model.User)

	fmt.Println(user)

	if !ok || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	userID := user.ID
	fmt.Println(userID)

	if err := c.BodyParser(&property); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given parameters are invalid",
		})
	}

	result, err := service.CreateProperty(app.DB, &property, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logrus.Error("Failed to get rows affected: ", err.Error())
	} else {
		logrus.Infof("Inserted property, %d row(s) affected", rowsAffected)
	}
	logrus.Infof("User ID from token: %d", userID)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":       "Accepted",
		"status":        "200",
		"Property Id":   property.ID,
		"Property Name": property.Name,
	})
}

func (app *App) DeletePropertyHandler(c *fiber.Ctx) error {

	id := c.Params("id")
	result_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Confusion",
			"status":  "Status Bad Request",
		})
	}

	if err := service.DeleteProperty(app.DB, result_id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given parameter is invalid",
			"status":  "Status Bad Request",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Deleted",
	})
}

func (app *App) UpdatePropertyHandler(c *fiber.Ctx) error {
	var property_update *model.PropertyUpdate

	if err := c.BodyParser(property_update); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given parameters are invalid",
		})
	}

	if err := service.UpdateProperty(app.DB, property_update); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Given parameters are invalid",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Accepted",
	})

}
