package controller

import (
	"justrun/database"
	"justrun/model/shoes"
	"net/http"
	"strconv"

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

func GetShoesTypeController(echoContext echo.Context) error {

	tipe, err := database.GetShoesType()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":     "Success",
		"shoes type": tipe,
	})
}

func ShoesbyTypeController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	shoes, err := database.ShoesByTypeID(id)
	tipe, err := database.ShoesTypeID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"type":   tipe,
		"shoe":   shoes,
	})
}
