package products

import (
	"log"
	"os"
	"path/filepath"

	"github.com/forge-io/backend/lib/models/product"
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

	err = db.AutoMigrate(&models.Product{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateProduct(product *models.Product) error {
	db := GetDB()

	return db.Create(product).Error
}
func GetAllProducts() ([]models.Product, error) {
	db := GetDB()

	var products []models.Product
	result := db.Find(&products)
	return products, result.Error
}

func GetProductByUUID(uuid string) (*models.Product, error) {
	db := GetDB()

	var product models.Product
	result := db.First(&product, "id = ?", uuid)
	return &product, result.Error
}

func UpdateProduct(uuid string, updatedData *models.Product) error {
	db := GetDB()

	var product models.Product
	if err := db.First(&product, "id = ?", uuid).Error; err != nil {
		return err
	}

	product.Model = updatedData.Model
	product.Category = updatedData.Category
	product.Year = updatedData.Year
	product.Brand = updatedData.Km
	product.Km = updatedData.Km
	product.Color = updatedData.Color
	product.Motor = updatedData.Motor
	product.Price = updatedData.Price
	product.Image = updatedData.Image



	return db.Save(&product).Error
}

func DeleteProduct(uuid string) error {
	db := GetDB()
	return db.Delete(&models.Product{}, "id = ?", uuid).Error
}