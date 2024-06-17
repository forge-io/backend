package api

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", healthcheck)

}

func healthcheck(c echo.Context) error {
	fmt.Print("healthCheck")
	return c.String(http.StatusOK, "OK")
}
