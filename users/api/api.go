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

	// user
	g.GET("/create", services.CreateUser)
}
