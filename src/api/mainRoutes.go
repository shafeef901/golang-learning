package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func MainRoutes(e *echo.Echo) {

	e.GET("login", handlers.Login)
	e.GET("/", handlers.Hello)
	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
}
