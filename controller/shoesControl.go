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

// func AddShoesController(c echo.Context) error {

// 	result, err := database.AddShoes()

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "hope all feeling well",
// 		"data":    respShoes.FromModelSlice(result),
// 	})
// }

// func AddShoesController(c echo.Context) error {

// 	shoes, err := database.AddShoes()

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Your post success!",
// 		"data":    shoes,
// 	})
// }
