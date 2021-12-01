package middleware

import (
	"justrun/config"
	"justrun/model/users"

	"github.com/labstack/echo/v4"
)

var db = config.DB

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {

	var db = config.DB
	var user users.Users

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}

	return true, nil
}
