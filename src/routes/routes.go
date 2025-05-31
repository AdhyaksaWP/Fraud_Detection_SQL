package routes

import (
	"database/sql"
	"fiber-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	userGroup := api.Group("/users")
	userGroup.Get("/", controllers.GetAllUser(db))
	userGroup.Post("/", controllers.InsertUser(db))
	userGroup.Delete("/:id", controllers.DeleteUser(db))

	cardGroup := api.Group("/cards")
	cardGroup.Get("/", controllers.GetAllCard(db))
	cardGroup.Post("/", controllers.InsertCard(db))
	cardGroup.Delete("/:card_number", controllers.DeleteCard(db))

	transactionGroup := api.Group("/transactions")
	transactionGroup.Get("/", controllers.GetAllTransactions(db))
	transactionGroup.Get("/:id", controllers.GetAllTransactionsByMerchantID(db))
	transactionGroup.Post("/", controllers.CreateTransaction(db))
}
