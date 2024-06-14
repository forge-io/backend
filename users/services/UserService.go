package services

import (
	"net/http"
	"time"

	handlers "github.com/forge-io/backend/lib/handlers/users"
	"github.com/forge-io/backend/lib/models"
	"github.com/labstack/echo/v4"
)

type UserCreationRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func CreateUser(c echo.Context) error {
	var userReq UserCreationRequest

	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body: "+err.Error())
	}

	u := models.User{
		Name:      userReq.Name,
		Age:       userReq.Age,
		Email:     userReq.Email,
		Password:  userReq.Password,
		Address:   userReq.Address,
		Phone:     userReq.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := handlers.CreateUser(&u)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func GetAllUsers(c echo.Context) error {
	users, err := handlers.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get users: "+err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByUUID(c echo.Context) error {
	uuid := c.Param("uuid")

	user, err := handlers.GetUserByUUID(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserByUUID(c echo.Context) error {
	uuid := c.Param("uuid")

	var userReq UserCreationRequest
	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body: "+err.Error())
	}

	updatedData := &models.User{
		Name:     userReq.Name,
		Age:      userReq.Age,
		Email:    userReq.Email,
		Password: userReq.Password,
		Phone:    userReq.Phone,
		Address:  userReq.Address,
	}

	err := handlers.UpdateUser(uuid, updatedData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user: "+err.Error())
	}

	return c.JSON(http.StatusOK, updatedData)
}

func DeleteUser(c echo.Context) error {
	uuid := c.Param("uuid")

	err := handlers.DeleteUser(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.NoContent(http.StatusNoContent)
}