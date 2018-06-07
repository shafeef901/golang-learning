package router

import (
	"api"
	"api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	////create groups
	g := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	////Setting middlewares

	//Always make sure that middlewares are USEd before the endpoints are defined
	middlewares.SetMainMiddleware(e)
	middlewares.SetAdminMiddlewares(g)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJWTMiddlewares(jwtGroup)

	////Setting group routes
	api.MainRoutes(e)
	api.AdminRoutes(g)
	api.CookieRoutes(cookieGroup)
	api.JWTRoutes(jwtGroup)

	return e
}
