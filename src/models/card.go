package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
)

type Card struct {
	CardNumber string `json:"card"`
	ID         int    `json:"id"`
}

func generateRandomCardNumber() string {
	length := rand.Intn(8) + 12

	cardNumber := ""
	for i := 0; i < length; i++ {
		cardNumber += strconv.Itoa(rand.Intn(10))
	}

	return cardNumber
}

func GetAllCard(db *sql.DB) ([]Card, error) {
	var cards []Card

	query := "SELECT * FROM credit_card"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var card Card
		if err := rows.Scan(&card.ID, &card.CardNumber); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func InsertCard(db *sql.DB, id int) (Card, error) {
	var card Card
	isCardNumberCreated := false
	cardNumber := ""

	// Check if ID is present in the card_holder table
	checkQuery := "SELECT id FROM card_holder WHERE id = $1"
	err := db.QueryRow(checkQuery, id).Scan(&card.ID)
	if err == sql.ErrNoRows {
		return card, fmt.Errorf("card_holder with id %d does not exist", id)
	} else if err != nil {
		return card, err
	}

	// Loop until the card number is unique before inserting it into the table
	for !isCardNumberCreated {
		cardNumber = generateRandomCardNumber()

		checkQuery := "SELECT card from credit_card WHERE card = $1"
		err := db.QueryRow(checkQuery, cardNumber).Scan(&card.ID)

		if err == sql.ErrNoRows {
			isCardNumberCreated = true
		} else if err != nil {
			return card, err
		}
	}

	insertQuery := "INSERT INTO credit_card (card, id_card_holder) VALUES ($1, $2) RETURNING id_card_holder"
	err = db.QueryRow(insertQuery, cardNumber, id).Scan(&card.ID)
	if err != nil {
		return card, err
	}

	card.CardNumber = cardNumber
	return card, nil
}

func DeleteCardByCardNumber(db *sql.DB, card_number string) error {
	query := "DELETE FROM credit_card WHERE card = $1"
	_, err := db.Exec(query, card_number)
	return err
}
