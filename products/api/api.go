package api

import (
	services "products/services"

	echo "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	// e.GET("/health-check", healthcheck.HealthCheck)

	e.POST("/create", services.CreateProduct)
	e.GET("/getall", services.GetAllProducts)
	e.GET("/getbyid/:uuid", services.GetProductByUUID)
	e.PUT("/update/:uuid", services.UpdateProductByUUID)
	e.DELETE("/delete/:uuid", services.DeleteProduct)
}
