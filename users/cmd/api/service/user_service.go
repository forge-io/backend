package service

import (
	"users/cmd/api/model"
	"errors"
)

var users = make(map[string]model.User)

func GetAllUsers() ([]model.User, error) {
	userList := []model.User{}
	for _, user := range users {
			userList = append(userList, user)
	}
	return userList, nil
}

func GetUserByID(id string) (*model.User, error) {
	user, exists := users[id]
	if !exists {
			return nil, nil
	}
	return &user, nil
}

func CreateUser(user model.User) (model.User, error) {
	if _, exists := users[user.ID]; exists {
			return model.User{}, errors.New("user already exists")
	}
	users[user.ID] = user
	return user, nil
}

func UpdateUser(id string, user model.User) (model.User, error) {
	if _, exists := users[id]; !exists {
			return model.User{}, errors.New("user not found")
	}
	users[id] = user
	return user, nil
}

func DeleteUser(id string) error {
	if _, exists := users[id]; !exists {
			return errors.New("user not found")
	}
	delete(users, id)
	return nil
}