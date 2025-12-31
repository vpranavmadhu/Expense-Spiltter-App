package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	return Config{
		DBUrl: dbUrl,
	}
}
