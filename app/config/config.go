package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	godotenv.Load() // 預設讀 .env
}

func GetConfig(key string) string {
	return os.Getenv(key)
}
