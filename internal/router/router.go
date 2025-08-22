package router

import (
	"database/sql"

	"github.com/doguhanniltextra/property_go/internal/handler"
	"github.com/doguhanniltextra/property_go/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(db *sql.DB) {
	app := fiber.New()

	Routers(app, db)

	app.Listen(":8080")
}

func Routers(app *fiber.App, db *sql.DB) {
	RouterUser(app, db)
	RouterProperty(app, db)
	RouterAdmin(app, db)
	RouterPropertyIncome(app, db)
}

func RouterUser(app *fiber.App, db *sql.DB) {
	api := app.Group("/users")

	appHandler := &handler.App{DB: db}

	api.Post("/register", middleware.RegisterValidation, middleware.CreateCookie, appHandler.CreateRegisterHandler)
	api.Post("/auth", middleware.ReadToken, appHandler.CreateAuthHandler)
}

func RouterProperty(app *fiber.App, db *sql.DB) {
	api := app.Group("/property")

	appHandler := &handler.App{DB: db}

	api.Post("/create", middleware.ReadToken, appHandler.CreatePropertyHandler)
	api.Delete("/delete/:id", middleware.ReadToken, appHandler.DeletePropertyHandler)
	api.Put("/update", middleware.ReadToken, appHandler.UpdatePropertyHandler)
}

func RouterAdmin(app *fiber.App, db *sql.DB) {
	api := app.Group("/admin")
	appHandler := &handler.App{DB: db}
	api.Get("/get-users", appHandler.GetAllUsersHandler)
}

func RouterPropertyIncome(app *fiber.App, db *sql.DB) {
	api := app.Group("/property-income")
	appHandler := &handler.App{DB: db}
	api.Post("/create", middleware.ReadToken, appHandler.CreatePropertyIncomeHandler)
}
