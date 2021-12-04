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

	sho := &shoes.Shoes{
		Name:        "Geist Ekiden",
		Price:       349000,
		BrandID:     4,
		ShoesTypeID: 1,
	}

	if err := config.DB.Create(&sho).Error; err != nil {
		fmt.Println("error")
	}
}

func TestAddShoes(t *testing.T) {

	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.POST("/api/shoes/add").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("status").String().Contains("success")
	result.Value("data").Array().NotEmpty()
}
