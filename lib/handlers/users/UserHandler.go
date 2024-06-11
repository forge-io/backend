package handlers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/forge-io/backend/lib/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	parentEnvPath, err := filepath.Abs(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalf("Error finding absolute path: %v", err)
	}

	// Load the parent .env file
	err = godotenv.Load(parentEnvPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("USER_DB")))

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateUser(user *models.User) error {
	db := GetDB()

	return db.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	db := GetDB()

	var user models.User
	result := db.First(&user, id)
	return &user, result.Error
}

func UpdateUser(user *models.User) error {
	db := GetDB()

	return db.Save(user).Error
}

func DeleteUser(id uint) error {
	db := GetDB()

	return db.Delete(&models.User{}, id).Error
}
