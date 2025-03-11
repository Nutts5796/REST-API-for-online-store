package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetProducts(c *fiber.Ctx, db *pgxpool.Pool) error {
	rows, err := db.Query(c.Context(), "SELECT id, name, description, price FROM products")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch products"})
	}

	products := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name, description string
		var price float64

		rows.Scan(&id, &name, &description, &price)

		product := map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
			"price":       price,
		}

		products = append(products, product)
	}

	return c.JSON(products)
}

func GetUsers(c *fiber.Ctx, db *pgxpool.Pool) error {
	rows, err := db.Query(context.Background(), "SELECT id, username, email FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить список пользователей",
		})
	}
	defer rows.Close()

	User := 0
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Ошибка при сканировании данных",
			})
		}
		users = append(users, user)
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}
