package handler

import (
	"database/sql"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/doguhanniltextra/property_go/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (app *App) GetAllUsersHandler(c *fiber.Ctx) error {
	rows, err := service.GetAllUsers(app.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		var id int64
		var name, email, password sql.NullString

		if err := rows.Scan(&id, &name, &email, &password); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error scanning users: " + err.Error(),
			})
		}

		u.ID = int(id)
		if name.Valid {
			u.Name = name.String
		}
		if email.Valid {
			u.Email = email.String
		}
		if password.Valid {
			u.Password = password.String
		}

		users = append(users, u)
	}

	return c.JSON(users)
}
