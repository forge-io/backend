package routes

import (
	"users/cmd/api/handler"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo){
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUserByID)
	e.POST("/users", handler.CreateUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
}