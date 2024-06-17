package services

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	jwtHandler "github.com/forge-io/backend/lib/handlers/jwtHandler"
	handlers "github.com/forge-io/backend/lib/handlers/users"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtCustomClaims struct {
	Email string `json:"name"`
	jwt.RegisteredClaims
}

func Authenticate(c echo.Context) error {
	parentEnvPath, err := filepath.Abs(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalf("Error finding absolute path: %v", err)
	}

	// Load the parent .env file
	err = godotenv.Load(parentEnvPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var authReq AuthRequest

	if err := c.Bind(&authReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body: "+err.Error())
	}

	valid := handlers.CheckUserPasswordHash(authReq.Email, authReq.Password)

	if valid != true {
		return c.JSON(echo.ErrUnauthorized.Code, valid)
	}

	jwt := jwtHandler.JwtWrapper{
		SecretKey:       os.Getenv("SECRET_KEY"),
		Issuer:          os.Getenv("ISSUER"),
		ExpirationHours: 24 * 365,
	}

	token, err := jwt.GenerateToken(authReq.Email)

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, valid)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
