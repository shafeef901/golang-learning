package handlers

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string "json:name"
	jwt.StandardClaims
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"joe",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	//Never sign the string barely like this; use some long string instead of mySecret from a file in the server
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	log.Println(token)
	return token, nil
}

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	pass := c.QueryParam("pass")

	if username == "joe" && pass == "secret" {

		//Test Cookie

		cookie := new(http.Cookie)
		cookie.Name = "SessionID"
		cookie.Value = "somehashedstring"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		//Cookie for token instead of header

		Tcookie := new(http.Cookie)
		Tcookie.Name = "JWTCookie"
		Tcookie.Value = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiam9lIiwiZXhwIjoxNTI4NDQ2MTQxLCJqdGkiOiJtYWluX3VzZXJfaWQifQ.YCCdxc3XVIAcgyCFrF76IwKljD30_1R9L4ukoC2mK-OvTSsA4IqLy_MkXnM8qNG6YrPS5-dSsTubEP0GQBlJhQ"
		Tcookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(Tcookie)

		//Creating JWT Token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Error creating Jwt token: ", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "Wrong Credentials")
}
