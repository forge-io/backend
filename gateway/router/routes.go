package router

import (
	"gateway/api"
	"gateway/api/groups"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	api.MainGroup(e)
	groups.ProductsGroup(e)
	groups.UserGroup(e)
	groups.AuthenticateGroup(e)

	return e
}
