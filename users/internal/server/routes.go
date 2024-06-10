package server

import (
	"net/http"

	handlers "github.com/forge-io/backend/lib/handlers/users"
	"github.com/forge-io/backend/lib/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)
	e.POST("/create", s.CreateUserController)

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) CreateUserController(c echo.Context) error {
	user := models.User{}

	handlers.CreateUser(&user)

	resp := user

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
