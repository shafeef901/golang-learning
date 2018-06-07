package handlers

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	log.Println("Username : ", claims["Name"], "UserID : ", claims["jti"])

	return c.String(http.StatusOK, "jwt main working")
}
