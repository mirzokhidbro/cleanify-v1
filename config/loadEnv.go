package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         string
	ServerPort     string
	DefaultOffset  string
	DefaultLimit   string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load(".env")

	if err != nil {
		return
	}

	config = Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBUserName:     os.Getenv("DB_USER"),
		DBUserPassword: os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBPort:         os.Getenv("DB_PORT"),
		ServerPort:     os.Getenv("PORT"),
		DefaultOffset:  os.Getenv(""),
		DefaultLimit: os.Getenv(""),
	}

	return
}
