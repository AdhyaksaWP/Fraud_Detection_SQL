package models

import (
	"database/sql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllUser(db *sql.DB) ([]User, error) {
	var users []User

	query := "SELECT * FROM card_holder"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func InsertUser(db *sql.DB, name string) (User, error) {
	var user User

	// Check if user already exists
	checkQuery := "SELECT id FROM card_holder WHERE name = $1"
	err := db.QueryRow(checkQuery, name).Scan(&user.ID)
	if err == nil {
		user.Name = name
		return user, nil
	}

	insertQuery := "INSERT INTO card_holder (name) VALUES ($1) RETURNING id"
	err = db.QueryRow(insertQuery, name).Scan(&user.ID)
	if err != nil {
		return user, err
	}

	user.Name = name
	return user, nil
}

func DeleteUserByID(db *sql.DB, id int) error {
	query := "DELETE FROM card_holder WHERE id = $1"
	_, err := db.Exec(query, id)
	return err
}
