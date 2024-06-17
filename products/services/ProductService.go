package services

import (
	"net/http"
	"time"

	handlers "github.com/forge-io/backend/lib/handlers/products"
	models "github.com/forge-io/backend/lib/models/product"
	"github.com/labstack/echo/v4"
)

type ProductCreationRequest struct {
	Model    string  `json:"model"`
	Category string  `json:"category"`
	Year     int     `json:"year"`
	Brand    string  `json:"brand"`
	Km       string  `json:"km"`
	Color    string  `json:"color"`
	Motor    string  `json:"motor"`
	Price    float64 `json:"price"`
	Image    string  `json:"image"`
}

func CreateProduct(c echo.Context) error {
	var productReq ProductCreationRequest

	if err := c.Bind(&productReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body: "+err.Error())
	}

	p := models.Product{
		Model:     productReq.Model,
		Category:  productReq.Category,
		Year:      productReq.Year,
		Brand:     productReq.Brand,
		Km:        productReq.Km,
		Color:     productReq.Color,
		Motor:     productReq.Motor,
		Price:     productReq.Price,
		Image:     productReq.Image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := handlers.CreateProduct(&p)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, p)
}

func GetAllProducts(c echo.Context) error {
	products, err := handlers.GetAllProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get product: "+err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func GetProductByUUID(c echo.Context) error {
	uuid := c.Param("uuid")

	product, err := handlers.GetProductByUUID(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}

func UpdateProductByUUID(c echo.Context) error {
	uuid := c.Param("uuid")

	var productReq ProductCreationRequest
	if err := c.Bind(&productReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body: "+err.Error())
	}

	updatedData := &models.Product{
		Model:     productReq.Model,
		Category:  productReq.Category,
		Year:      productReq.Year,
		Brand:     productReq.Brand,
		Km:        productReq.Km,
		Color:     productReq.Color,
		Motor:     productReq.Motor,
		Price:     productReq.Price,
		Image:     productReq.Image,
		DeletedAt: time.Now(),
	}

	err := handlers.UpdateProduct(uuid, updatedData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update product: "+err.Error())
	}

	return c.JSON(http.StatusOK, updatedData)
}

func DeleteProduct(c echo.Context) error {
	uuid := c.Param("uuid")

	err := handlers.DeleteProduct(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	return c.NoContent(http.StatusNoContent)
}
