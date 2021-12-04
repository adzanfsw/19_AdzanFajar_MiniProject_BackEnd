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

	tipe := &shoes.ShoesType{
		Type: "Marathon",
	}

	if err := config.DB.Create(&tipe).Error; err != nil {
		fmt.Println("error")
	}
}

func TestGetShoesType(t *testing.T) {

	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/api/shoes-type").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("status").String().Contains("Success")
	result.Value("shoes type").Array().NotEmpty()
}
