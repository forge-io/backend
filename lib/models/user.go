package models

import (

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
}
