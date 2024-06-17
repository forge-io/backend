package handlers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/forge-io/backend/lib/models/user"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckUserPasswordHash(email string, password string) bool {
	user, err := GetUserByEmail(email)

	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

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

	uPass := user.Password

	hash, err := HashPassword(uPass)

	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	user.Password = hash

	return db.Create(user).Error
}

func GetAllUsers() ([]models.User, error) {
	db := GetDB()

	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByUUID(uuid string) (*models.User, error) {
	db := GetDB()

	var user models.User
	result := db.First(&user, "id = ?", uuid)
	return &user, result.Error
}

func GetUserByEmail(email string) (*models.User, error) {
	db := GetDB()

	var user models.User
	result := db.First(&user, "email = ?", email)
	return &user, result.Error
}

func UpdateUser(uuid string, updatedData *models.User) error {
	db := GetDB()

	var user models.User
	if err := db.First(&user, "id = ?", uuid).Error; err != nil {
		return err
	}

	user.Name = updatedData.Name
	user.Age = updatedData.Age
	user.Email = updatedData.Email
	user.Password = updatedData.Password
	user.Phone = updatedData.Phone
	user.Address = updatedData.Address

	return db.Save(&user).Error
}

func DeleteUser(uuid string) error {
	db := GetDB()
	return db.Delete(&models.User{}, "id = ?", uuid).Error
}
