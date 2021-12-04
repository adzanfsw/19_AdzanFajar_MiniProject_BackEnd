package test

import (
	"fmt"
	"justrun/config"
	"justrun/model/review"
	"net/http"
	"testing"
)

func init() {
	config.InitDB("justrun_test")

	rev := &review.Review{
		Rating:  3,
		Review:  "Bla bla bla",
		UserID:  3,
		ShoesID: 1,
	}

	if err := config.DB.Create(&rev).Error; err != nil {
		fmt.Println("error")
	}
}

func TestUpdateReview(t *testing.T) {

	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.PUT("/api/reviews/update/:id").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("status").String().Contains("Update Success")
	result.Value("data").Array().NotEmpty()
}
