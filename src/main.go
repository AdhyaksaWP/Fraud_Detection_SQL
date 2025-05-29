package main

import (
	"database/sql"
	"fiber-app/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv for DB DSN
	godotenv.Load(".env")
	dsn := os.Getenv("DB_DSN")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to open Postgres DB:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping Postgres DB:", err)
	}

	fmt.Println("Connected to Postgres!")

	app := fiber.New()

	// Pass in the main fiber app and the DB's DSN
	routes.SetupRoutes(app, db)

	app.Listen(":3000")
}
