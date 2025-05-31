package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type Transaction struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Amount      float64   `json:"amount"`
	CardNumber  string    `json:"card_number"`
	ID_Merchant int       `json:"id_merchant"`
}

func GetAllTransactions(db *sql.DB) ([]Transaction, error) {
	var transactions []Transaction

	query := "SELECT * FROM transaction"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.ID,
			&transaction.Date, &transaction.Amount,
			&transaction.CardNumber, &transaction.ID_Merchant); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetAllTransactionsByMerchantsID(db *sql.DB, id_merchant int) ([]Transaction, error) {
	var transactions []Transaction

	query := "SELECT * FROM transaction WHERE id_merchant = $1"
	rows, err := db.Query(query, id_merchant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.ID,
			&transaction.Date, &transaction.Amount,
			&transaction.CardNumber, &transaction.ID_Merchant); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func CreateTransaction(db *sql.DB, amount float64, card string, id_merchant int) (Transaction, error) {
	var transaction Transaction

	// Validate card exists
	checkCardQuery := "SELECT card FROM credit_card where card = $1"
	err := db.QueryRow(checkCardQuery, card).Scan(&transaction.CardNumber)
	if err == sql.ErrNoRows {
		return transaction, fmt.Errorf("credit card with number %s does not exist", card)
	} else if err != nil {
		return transaction, err
	}

	checkIDMerchantQuery := "SELECT id FROM merchant where id = $1"
	err = db.QueryRow(checkIDMerchantQuery, id_merchant).Scan(&transaction.ID_Merchant)
	if err == sql.ErrNoRows {
		return transaction, fmt.Errorf("merchantg with id %d does not exist", id_merchant)
	} else if err != nil {
		return transaction, err
	}

	id := 0
	isIDCreated := false

	for !isIDCreated {
		id = rand.Intn(100000)

		checkQuery := `SELECT * FROM transaction WHERE id = $1`
		err := db.QueryRow(checkQuery, id).Scan(&transaction.ID)

		if err == sql.ErrNoRows {
			isIDCreated = true
		} else if err != nil {
			return transaction, err
		}
	}

	now := time.Now()
	query := `
		INSERT INTO transaction (id, date, amount, card, id_merchant)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = db.Exec(query, id, now.Format("2006-01-02 15:04:05"), amount, card, id_merchant)
	if err != nil {
		return transaction, err
	}

	transaction.ID = id
	transaction.Date = now
	transaction.Amount = amount
	transaction.CardNumber = card
	transaction.ID_Merchant = id_merchant

	return transaction, nil
}
