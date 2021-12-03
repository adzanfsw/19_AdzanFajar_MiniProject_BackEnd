package controller

import (
	"justrun/database"
	"justrun/model/shoes"
	"net/http"
	"strconv"

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

func GetShoesBrandController(echoContext echo.Context) error {

	merk, err := database.GetShoesBrand()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":      "Success",
		"shoes brand": merk,
	})
}

func ShoesbyBrandController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	shoes, err := database.ShoesByBrandID(id)
	brand, err := database.ShoesBrandID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"brand":  brand,
		"shoe":   shoes,
		"status": "Success",
	})
}
