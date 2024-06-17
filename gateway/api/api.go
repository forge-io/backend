package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	handlers "github.com/forge-io/backend/lib/handlers/jwtHandler"
	echo "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", healthcheck)

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8081",
	})

	e.Any("/users", echo.WrapHandler(authenticate(http.HandlerFunc(proxy.ServeHTTP))))
}

func healthcheck(c echo.Context) error {
	fmt.Print("healthCheck")
	return c.String(http.StatusOK, "OK")
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		token := strings.Split(auth, "Bearer ")

		jwt := handlers.JwtWrapper{
			SecretKey:       "macaco-prego-123",
			Issuer:          "go-grpc-auth-svc",
			ExpirationHours: 24 * 365,
		}

		_, err := jwt.ValidateToken(token[1])

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

func proxy(path string, target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		targetURL := target + r.URL.Path
		req, err := http.NewRequest(r.Method, targetURL, r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		req.Header = r.Header

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(resp.StatusCode)

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}
}
