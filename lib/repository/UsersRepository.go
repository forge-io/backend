package repository

import (
	models "github.com/forge-io/backend/lib/models/user"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}
