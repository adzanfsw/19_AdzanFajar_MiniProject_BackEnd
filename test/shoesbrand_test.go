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

	brand := &shoes.ShoesBrand{
		Brand: "Adidas",
	}

	if err := config.DB.Create(&brand).Error; err != nil {
		fmt.Println("error")
	}
}

func TestGetShoesBrand(t *testing.T) {

	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/api/shoes-brand").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("status").String().Contains("Success")
	result.Value("shoes brand").Array().NotEmpty()
}
