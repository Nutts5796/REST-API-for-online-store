package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(app *fiber.App, db *pgxpool.Pool) {

	app.Get("/api/users", func(c *fiber.Ctx) error {
		return GetUsers(c, db)
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return Register(c, db)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return Login(c, db)
	})

	app.Post("/logout", Logout)

	app.Post("/refresh-token", RefreshToken)

}
