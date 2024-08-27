package models

import (
	"backend/config"
	"database/sql"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func FetchAllUsers() ([]User, error) {
	var users []User

	rows, err := config.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FetchUserByID(id int) (*User, error) {
	var user User

	row := config.DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
