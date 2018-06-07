package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func CookieRoutes(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}
