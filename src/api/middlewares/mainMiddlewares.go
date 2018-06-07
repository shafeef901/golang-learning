package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//This package is for middlewares on echo instance; e

func SetMainMiddleware(e *echo.Echo) {
	//Use middleware below to load website static content from the static folder; static folder contains html,css,js
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static",
	}))

	e.Use(serverHeader)
}

//Custom Middleware
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Accelerate")
		c.Response().Header().Set("Test Header", "Header Value")

		return next(c)
	}
}
