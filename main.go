package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nutts5796/REST-API-for-online-store/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	dbpool, err := pgxpool.New(nil, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	app := fiber.New()

	api := app.Group("/api")

	handler.RegisterRoutes(api, dbpool)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
