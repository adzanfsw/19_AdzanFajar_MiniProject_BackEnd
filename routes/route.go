package routes

import (
	"justrun/factory"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	presenter := factory.Init()

	//initiate
	e := echo.New()

	e.POST("/articles", presenter.ShoesPresentation.AddShoes)

	return e
}
