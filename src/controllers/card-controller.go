package controllers

import (
	"database/sql"
	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllCard(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cards, err := models.GetAllCard(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch credit_card table",
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(cards)
	}
}

func InsertCard(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Card

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Invalid request body",
				"details": err.Error(),
			})
		}

		card, err := models.InsertCard(db, input.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to insert card",
				"details": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(card)
	}
}

func DeleteCard(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cardNumberStr := c.Params("card_number")

		err := models.DeleteCardByCardNumber(db, cardNumberStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete card",
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
