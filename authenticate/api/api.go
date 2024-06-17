package api

import (
	services "authenticate/services"

	echo "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	// e.GET("/health-check", healthcheck.HealthCheck)

	e.POST("/authenticate", services.Authenticate)
}
