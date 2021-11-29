package controller

import (
	"justrun/database"
	"justrun/model/users"
	"net/http"

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
		"status": "success",
		"data":   result,
	})
}
