package test

import (
	"fmt"
	"justrun/config"
	"justrun/model/shoes"
	"net/http"
	"testing"
)

func init() {
	config.InitDB("justrun_test")

	desc := &shoes.ShoesDescription{
		ShoesID:     11,
		Description: "Sepatu ini mantap!!",
		Color:       "Red",
		PurchaseURL: "https://910.id/reseller",
	}

	if err := config.DB.Create(&desc).Error; err != nil {
		fmt.Println("error")
	}
}

func TestGetShoesDesc(t *testing.T) {

	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/api/shoes-desc").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("status").String().Contains("Success")
	result.Value("shoes desc").Array().NotEmpty()
}
