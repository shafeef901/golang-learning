package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func JWTRoutes(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
