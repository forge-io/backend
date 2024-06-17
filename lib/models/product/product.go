package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Model     string    `json:"model"`
	Category  string    `json:"category"`
	Year      int       `json:"year"`
	Brand     string    `json:"brand"`
	Km        string    `json:"km"`
	Color     string    `json:"color"`
	Motor     string    `json:"motor"`
	Price     float64   `json:"price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
