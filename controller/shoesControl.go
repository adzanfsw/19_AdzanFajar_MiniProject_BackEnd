package controller

import (
	"net/http"

	"justrun/database"
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

func GetShoesController(echoContext echo.Context) error {

	shoes, err := database.GetShoes()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"shoes":  shoes,
	})
}
