package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURL string
	Port     string
	DBName   string
}

var AppConfig *Config

func New() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Print(err)
	}
	AppConfig = &Config{
		MongoURL: os.Getenv("MONGO_URL"),
		Port:     getenv("PORT", "3000").(string),
		DBName:   os.Getenv("DB_NAME"),
	}
	return AppConfig
}

func getenv(key string, fallback interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
