package models

import (
	"backend/config"
	"backend/controllers"
)

type User struct{
	ID string `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	}

func FetchAllUsers(*User, error)  {
	var users []User

	row, err:= config.DB.Query("SELECT id, username, email FROM users")
	if err!=nil {
		return nil, err
	}

	defer row.Close()

	for row.Next(){
		var user User
		if err:= row.Scan(&user.ID, &user.Username, &user.Email)
	}
}