package handler

import (
	"net/http"
	"users/cmd/api/model"
	"users/cmd/api/service"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := service.GetAllUsers()
	if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := service.GetUserByID(id)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
	}
	if user == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
	}

	newUser, err := service.CreateUser(user)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User
	if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
	}

	updatedUser, err := service.UpdateUser(id, user)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := service.DeleteUser(id)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}