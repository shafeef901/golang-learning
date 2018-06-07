package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("I'm in")
		cookie, err := c.Cookie("SessionID")

		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you don't have any cookie")
			}
			log.Println(err)
			return err
		}

		//check from DB
		if cookie.Value == "somehashedstring" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Wrong Cookie")
	}
}
