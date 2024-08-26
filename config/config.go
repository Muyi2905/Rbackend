package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	connStr := os.Getenv("POSTGRES_URL")
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return DB.Ping()
}
