package controller

import (
	"justrun/database"
	"justrun/model/wishlist"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddWishController(echoContext echo.Context) error {

	var wishReq wishlist.Wishlist
	echoContext.Bind(&wishReq)

	result, err := database.AddWishlist(wishReq)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"data":   result,
		"status": "Success",
	})
}
