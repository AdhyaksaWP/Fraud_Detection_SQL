package routes

import (
	"database/sql"
	"fiber-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	userGroup := api.Group("/users")
	userGroup.Post("/", controllers.InsertUser(db))
	userGroup.Get("/", controllers.GetAllUser(db))
	userGroup.Delete("/:id", controllers.DeleteUser(db))
}
