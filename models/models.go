package models

type Register struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json"password "`
	UserName string `json:"username"`
}
type Login struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json"password "`
}
