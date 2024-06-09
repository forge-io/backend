package main

import (
	"users/cmd/api/config"
	"users/internal/routes"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config.Init()

	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
