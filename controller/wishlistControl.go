package controller

import (
	"justrun/database"
	"justrun/model/wishlist"
	"net/http"
	"strconv"

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

func GetWishController(echoContext echo.Context) error {

	wish, err := database.GetWishlists()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":    "Success",
		"wishlists": wish,
	})
}

func GetWishbyUserIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	wish, err := database.WishByUserID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":    "Success",
		"wishlists": wish,
	})
}

func GetWishbyShoesIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	wish, err := database.WishByShoesID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":    "Success",
		"wishlists": wish,
	})
}
