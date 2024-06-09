package config

import (
	"log"
	"os"

	// Substitua 'project-root' pelo caminho correto do seu módulo
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
			log.Fatalf("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}