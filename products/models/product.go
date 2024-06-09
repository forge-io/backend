package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4()"`
	Model     string         `json:"model"`
	Category  string         `json:"category"`
	year      int            `json:"year"`
	Brand     string         `json:"brand"`
	Km        string         `json:"km"`
	Collor    string         `json:"collor"`
	Motor     string         `json:"motor"`
	Price     float64        `json:"price"`
	registerProduct time.Time`json:"registerProduct"`
	buy       time.Time      `json:"buy"`
}
