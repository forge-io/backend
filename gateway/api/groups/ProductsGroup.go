package groups

import (
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	handlers "github.com/forge-io/backend/lib/handlers/jwtHandler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func ProductsGroup(e *echo.Echo) {
	parentEnvPath, err := filepath.Abs(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalf("Error finding absolute path: %v", err)
	}

	// Load the parent .env file
	err = godotenv.Load(parentEnvPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	url, _ := url.Parse("http://localhost:" + os.Getenv("PRODUCTS_PORT"))
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Wrap the proxy middleware with the authentication middleware
	productsGroup := e.Group("/products")
	productsGroup.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {

			req := context.Request()
			res := context.Response().Writer

			//may be some extra validations before sending request like Auth etc.
			auth := req.Header.Get("Authorization")
			token := strings.Split(auth, "Bearer ")

			if auth == "" {
				return context.JSON(echo.ErrUnauthorized.Code, map[string]string{"message": "Invalid token"})
			}

			jwt := handlers.JwtWrapper{
				SecretKey:       os.Getenv("SECRET_KEY"),
				Issuer:          os.Getenv("ISSUER"),
				ExpirationHours: 24 * 365,
			}

			_, err := jwt.ValidateToken(token[1])
			if err != nil {
				return context.JSON(echo.ErrUnauthorized.Code, map[string]string{"message": "Invalid token"})
			}

			// Update the headers to allow for SSL redirection
			req.Host = url.Host
			req.URL.Host = url.Host
			req.URL.Scheme = url.Scheme

			//trim reverseProxyRoutePrefix
			path := req.URL.Path
			req.URL.Path = strings.TrimPrefix(path, "/products")

			fmt.Print("Proxing request to gateway -> ", url.String()+path)

			// ServeHttp is non blocking and uses a go routine under the hood
			proxy.ServeHTTP(res, req)
			return nil
		}
	})
}
