package bussiness

import (
	"justrun/features/shoes"

	"github.com/go-playground/validator/v10"
)

type shoeCase struct {
	shoesData shoes.Data
	validate  *validator.Validate
}

func NewShoeBussiness(shoe shoes.Data) shoes.Bussiness {

	return &shoeCase{
		shoesData: shoe,
		validate:  validator.New(),
	}
}

func (sc *shoeCase) CreateData(data shoes.Shoes) (resp shoes.Shoes, err error) {

	if err := sc.validate.Struct(data); err != nil {
		return shoes.Shoes{}, err
	}

	resp, err = sc.shoesData.InsertData(data)
	if err != nil {
		return shoes.Shoes{}, err
	}
	return shoes.Shoes{}, nil
}

func (sc *shoeCase) GetAllData(search string) (resp []shoes.Shoes) {

	resp = sc.shoesData.SelectData(search)
	return
}
