package controller

import (
	"justrun/database"
	"justrun/model/shoes"
	"net/http"
	"strconv"

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

func GetShoesDescController(echoContext echo.Context) error {

	desc, err := database.GetShoesDesc()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":     "Success",
		"shoes desc": desc,
	})
}

func DescbyShoesIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	shoes, err := database.ShoesByID(id)
	desc, err := database.DescByShoesID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":           "Success",
		"shoe description": desc,
		"shoe":             shoes,
	})
}
