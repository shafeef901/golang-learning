package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func AdminRoutes(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
