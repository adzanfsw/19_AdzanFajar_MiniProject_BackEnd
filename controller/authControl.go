package controller

import (
	"justrun/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthLogin(c echo.Context) error {

	username := c.QueryParam("username")
	password := c.QueryParam("password")

	token := ""

	if username == "admin" && password == "admin" {
		var err error
		token, err = middleware.CreateToken(1)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Success",
		"JWT Token": token,
	})
}
