package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "cookie main working")
}
