package controller

import (
	"net/http"

	"justrun/library/database"
	"justrun/model/shoes"

	"github.com/labstack/echo/v4"
)

func AddShoesController(echoContext echo.Context) error {

	var shoeReq shoes.Shoes
	echoContext.Bind(&shoeReq)

	result, err := database.AddShoes(shoeReq)
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
