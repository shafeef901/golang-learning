package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Cat struct {
	Name string "json:name"
	Type string "json:type"
}

func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your cat name: %s \n Your cat type is: %s \n", catName, catType))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Enter a data type",
	})

}

func AddCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed Reading Request Body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(body, &cat)
	if err != nil {
		log.Printf("Failed Unmarshalling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("your cat: %#v", cat)
	return c.String(http.StatusOK, "this is your cat")
}
