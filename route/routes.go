package route

import (
	"justrun/controller"

	"github.com/labstack/echo/v4"
)

func RouteShoes() *echo.Echo {

	e := echo.New()

	e.POST("/api/shoes/add", controller.AddShoesController)
	e.POST("/api/shoes-type/add", controller.AddShoesTypeController)
	e.POST("/api/shoes-brand/add", controller.AddShoesBrandController)
	e.POST("/api/shoes-desc/add", controller.AddShoesDescController)

	return e
}
