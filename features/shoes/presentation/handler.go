package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"justrun/features/shoes"
	"justrun/features/shoes/presentation/request"
	"justrun/features/shoes/presentation/response"
)

type ShoesHandler struct {
	shoesBusiness shoes.Bussiness
}

func AddShoesHandler(sb shoes.Bussiness) *ShoesHandler {
	return &ShoesHandler{
		shoesBusiness: sb,
	}
}

func (sh *ShoesHandler) AddShoes(c echo.Context) error {

	var newShoes request.Shoes

	c.Bind(&newShoes)

	resp, err := sh.shoesBusiness.CreateData(request.ToShoes(newShoes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    response.FromShoes(resp),
	})
}
