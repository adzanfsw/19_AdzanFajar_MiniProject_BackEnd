package controller

import (
	"net/http"
	"strconv"

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

func ShoesbyIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	shoes, err := database.ShoesByID(id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"shoe":   shoes,
	})
}

func UpdateShoesController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	var shoeReq shoes.Shoes
	echoContext.Bind(&shoeReq)

	result, err := database.UpdateShoes(id, shoeReq)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Update Success",
		"data":   result,
	})
}

func DeleteShoesController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	shoes, err := database.DeleteShoes(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   shoes,
	})
}
