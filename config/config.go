package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURL string
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Print(err)
	}
	return &Config{
		MongoURL: os.Getenv("MONGO_URL"),
	}
}
