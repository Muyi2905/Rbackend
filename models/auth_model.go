package models

type Register struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password "`
	UserName  string `json:"username"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password "`
}
