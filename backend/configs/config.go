package configs

import (
	"fmt"
)

type Config struct {
	DBUrl string
}

func LoadConfig() Config {
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASS")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")
	host := "localhost"
	user := "testuser"
	password := "testpass"
	dbname := "esdb"
	port := "5432"

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	return Config{
		DBUrl: dbUrl,
	}
}
