package api

import (
	services "users/services"

	echo "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	// e.GET("/health-check", healthcheck.HealthCheck)

	e.POST("/create", services.CreateUser)
	e.GET("/getall", services.GetAllUsers)
	e.GET("/getbyid/:uuid", services.GetUserByUUID)
	e.PUT("/update/:uuid", services.UpdateUserByUUID)
	e.DELETE("/delete/:uuid", services.DeleteUser)
}
