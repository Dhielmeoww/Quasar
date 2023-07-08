// config/config.go

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName  string
	AppEnv   string
	AppPort  string
	DbHost   string
	DbDriver string
	DbName   string
	DbUser   string
	DbPass   string
	DbPort   string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	return Config{
		AppName:  os.Getenv("APP_NAME"),
		AppEnv:   os.Getenv("APP_ENV"),
		AppPort:  os.Getenv("APP_PORT"),
		DbHost:   os.Getenv("MYSQL_DB_HOST"),
		DbDriver: os.Getenv("MYSQL_DB_DRIVER"),
		DbName:   os.Getenv("MYSQL_DB_NAME"),
		DbUser:   os.Getenv("MYSQL_DB_USER"),
		DbPass:   os.Getenv("MYSQL_DB_PASS"),
		DbPort:   os.Getenv("MYSQL_DB_PORT"),
	}
}
