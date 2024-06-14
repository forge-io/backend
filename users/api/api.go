package api

import (
	services "users/services"

	echo "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	// e.GET("/health-check", healthcheck.HealthCheck)

}

func UserGroup(e *echo.Echo) {
	g := e.Group("/user")

	g.POST("/create", services.CreateUser)
	g.GET("/getall", services.GetAllUsers)
	g.GET("/getbyid/:uuid", services.GetUserByUUID)
	g.PUT("/update/:uuid", services.UpdateUserByUUID)
	g.DELETE("/delete/:uuid", services.DeleteUser)
}
