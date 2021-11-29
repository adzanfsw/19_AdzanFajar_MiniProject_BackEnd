package controller

import (
	"justrun/database"
	"justrun/model/shoes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddShoesTypeController(echoContext echo.Context) error {

	var typeReq shoes.ShoesType
	echoContext.Bind(&typeReq)

	result, err := database.AddShoesType(typeReq)
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
