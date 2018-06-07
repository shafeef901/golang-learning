package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetJWTMiddlewares(g *echo.Group) {
	//Looking Up for token in the cookie; By default, it checks in header under AuthScheme Bearer; Checl JWTConfig docs
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		TokenLookup:   "cookie:JWTCookie",
	}))
}
