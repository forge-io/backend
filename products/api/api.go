package api

import (
	services "products/services"

	echo "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	// e.GET("/health-check", healthcheck.HealthCheck)

}

func UserGroup(e *echo.Echo) {
	g := e.Group("/product")

	g.POST("/create", services.CreateProduct)
	g.GET("/getall", services.GetAllProducts)
	g.GET("/getbyid/:uuid", services.GetProductByUUID)
	g.PUT("/update/:uuid", services.UpdateProductByUUID)
	g.DELETE("/delete/:uuid", services.DeleteProduct)
}
