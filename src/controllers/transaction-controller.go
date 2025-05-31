package controllers

import (
	"database/sql"
	"fiber-app/models"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func GetAllTransactions(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		transactions, err := models.GetAllTransactions(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch transactions",
			})
		}

		return c.Status(fiber.StatusOK).JSON(transactions)
	}
}

func GetAllTransactionsByMerchantID(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid merchant ID",
			})
		}

		transactions, err := models.GetAllTransactionsByMerchantsID(db, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch transactions for merchant",
			})
		}

		return c.Status(fiber.StatusOK).JSON(transactions)
	}
}

func CreateTransaction(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Transaction
		pythonPath := os.Getenv("PYTHON_PATH")

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid request body",
				"details": err.Error(),
			})
		}

		cmd := exec.Command(pythonPath, "./controllers/predictor.py", fmt.Sprintf("%.2f", input.Amount))
		output, err := cmd.CombinedOutput() // Capture stdout and stderr
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to run python script",
				"details": string(output), // send back the real Python error
			})
		}

		result := string(output)
		if result == "1\r\n" { // Python script outputs 1 for fraud
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Fraudulent transaction detected!",
			})
		}

		transaction, err := models.CreateTransaction(db, input.Amount, input.CardNumber, input.ID_Merchant)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to create transaction",
				"details": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(transaction)
	}
}
