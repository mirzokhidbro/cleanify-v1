package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DBHost         string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         int
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
		DBHost:         cast.ToString(os.Getenv("DB_HOST")),
		DBUserName:     cast.ToString(os.Getenv("DB_USER")),
		DBUserPassword: cast.ToString(os.Getenv("DB_PASSWORD")),
		DBName:         cast.ToString(os.Getenv("DB_NAME")),
		DBPort:         cast.ToInt(os.Getenv("DB_PORT")),
		ServerPort:     cast.ToString(os.Getenv("PORT")),
		DefaultOffset:  cast.ToString(os.Getenv("")),
		DefaultLimit:   cast.ToString(os.Getenv("")),
	}

	return
}
