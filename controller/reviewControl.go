package controller

import (
	"justrun/database"
	"justrun/model/review"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddReviewController(echoContext echo.Context) error {

	var revReq review.Review
	echoContext.Bind(&revReq)

	result, err := database.AddReview(revReq)
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

func GetReviewsController(echoContext echo.Context) error {

	rev, err := database.GetReviews()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"reviews": rev,
	})
}

func UpdateReviewController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	var reviewReq review.Review
	echoContext.Bind(&reviewReq)

	result, err := database.UpdateReview(id, reviewReq)
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

func GetReviewsIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	re, err := database.ReviewByID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"review": re,
	})
}

func GetReviewsbyShoesIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	re, err := database.ReviewByShoesID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"review": re,
	})
}

func GetReviewsbyRatingController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	re, err := database.ReviewByRating(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"review": re,
	})
}

func GetReviewsbyUserIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	re, err := database.ReviewByUserID(id)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"review": re,
	})
}
