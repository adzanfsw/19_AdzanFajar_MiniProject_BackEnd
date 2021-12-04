package test

import (
	"fmt"
	"justrun/config"
	"justrun/model/wishlist"
	"net/http"
	"testing"
)

func init() {
	config.InitDB("justrun_test")

	wish := &wishlist.Wishlist{
		UserID:  1,
		ShoesID: 3,
	}

	if err := config.DB.Create(&wish).Error; err != nil {
		fmt.Println("error")
	}
}

func TestGetWishlist(t *testing.T) {

	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/api/wishlists").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("status").String().Contains("Success")
	result.Value("wishlists").Array().NotEmpty()
}
