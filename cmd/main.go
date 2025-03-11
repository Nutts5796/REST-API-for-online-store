package main

import (
	"context"
	"log"

	"github.com/Nutts5796/REST-API-for-online-store/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	app := fiber.New()

	db, err := pgxpool.New(context.Background(), "postgres://postgres:password@db:5432/REST-API-for-online-store")
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	handler.RegisterRoutes(app, db)

	log.Println("Сервер запущен на :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
