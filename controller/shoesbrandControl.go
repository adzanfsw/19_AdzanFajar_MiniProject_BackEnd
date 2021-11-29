package controller

import (
	"justrun/database"
	"justrun/model/shoes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddShoesBrandController(echoContext echo.Context) error {

	var brandReq shoes.ShoesBrand
	echoContext.Bind(&brandReq)

	result, err := database.AddShoesBrand(brandReq)
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
