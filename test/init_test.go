package test

import (
	"justrun/config"
	"justrun/route"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
)

var (
	echoHandler *echo.Echo
	server      *httptest.Server
)

func init() {
	config.InitDB("justrun_test")
	config.AutoMigrate()

	echoHandler = route.RouteShoes()
	server = httptest.NewServer(echoHandler)
}

func TearUp() func() {

	return func() {
		config.InitDB("justrun_test")
		config.DB.Exec("TRUNCATE TABLE shoes_brands;")
		config.DB.Exec("TRUNCATE TABLE shoes_descriptions;")
		config.DB.Exec("TRUNCATE TABLE shoes_types;")
		config.DB.Exec("TRUNCATE TABLE wishlists;")
		config.DB.Exec("TRUNCATE TABLE shoes;")
		config.DB.Exec("TRUNCATE TABLE users;")
		config.DB.Exec("TRUNCATE TABLE reviews;")
	}
}

// GetHttpExpect Get cached expect, create new if nil
func GetHTTPExpect(t *testing.T) *httpexpect.Expect {
	if server == nil {
		server = httptest.NewServer(echoHandler)
	}
	return NewHTTPExpect(t)
}

func NewHTTPExpect(t *testing.T) *httpexpect.Expect {
	// TODO : printer set, only if the testing is failed
	// https://github.com/gavv/httpexpect/issues/69
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: server.URL,
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
		Reporter: t,
	})
}
