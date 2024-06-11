package usecase

import (
	"github.com/forge-io/backend/lib/models"
	"github.com/forge-io/backend/lib/repository"
)

type UserUsecase struct{
		repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase{
	return UserUsecase{
		repository: repo,
	}
}

func (uu *UserUsecase) GetAllUsers() ([]models.User, error){
	return uu.repository.GetAllUsers()
}

func (uu *UserUsecase) CreateUser(user models.User) (models.User, error){
	userId, err := uu.repository.CreateUser(user)
	if err != nil{
		return models.User{}, err
	}
	user.ID = userId

	return user, nil
}