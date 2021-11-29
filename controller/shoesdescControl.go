package controller

import (
	"justrun/database"
	"justrun/model/shoes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddShoesDescController(echoContext echo.Context) error {

	var descReq shoes.ShoesDescription
	echoContext.Bind(&descReq)

	result, err := database.AddShoesDescription(descReq)
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
