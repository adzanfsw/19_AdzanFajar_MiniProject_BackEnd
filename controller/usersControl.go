package controller

import (
	"justrun/database"
	"justrun/model/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddUsersController(echoContext echo.Context) error {

	var userReq users.Users
	echoContext.Bind(&userReq)

	result, err := database.AddUsers(userReq)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Insert User Success",
		"data":   result,
	})
}

func GetUsersController(echoContext echo.Context) error {

	users, err := database.GetUsers()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"users":  users,
	})
}

func UserbyIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	users, err := database.UserByID(id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"user":   users,
	})
}
