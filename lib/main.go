package main

import (
	"time"

	handlers "github.com/forge-io/backend/lib/handlers/users"
	"github.com/forge-io/backend/lib/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	user := models.User{
		ID:        uuid.New(),
		Name:      "teste",
		Age:       19,
		Email:     "adriano.molin@hotmail.com",
		Password:  "pass",
		Phone:     "48996915228",
		Address:   "teste",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	handlers.CreateUser(&user)
}
