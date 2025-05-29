package controllers

import (
	"database/sql"
	"fiber-app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := models.GetAllUser(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch card_holder table",
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(users)
	}
}

func InsertUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.User

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid request body",
				"details": err.Error(),
			})
		}

		user, err := models.InsertUser(db, input.Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to insert user",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

func DeleteUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		err = models.DeleteUserByID(db, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete user",
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
