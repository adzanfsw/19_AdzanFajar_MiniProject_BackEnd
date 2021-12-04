package test

import (
	"fmt"
	"justrun/config"
	"justrun/model/users"
	"net/http"
	"testing"
)

func init() {
	config.InitDB("justrun_test")

	user := &users.Users{
		FirstName: "Adzan",
		LastName:  "Fajar",
		Gender:    "Laki-laki",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		fmt.Println("error")
	}
}

func TestUpdateUser(t *testing.T) {

	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.DELETE("/api/users/delete/:id").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("status").String().Contains("Delete Success")
	result.Value("data").Array().NotEmpty()
}
