package database

import (
	"justrun/config"
	shoe "justrun/model/shoes"
)

func AddShoes(shoes shoe.Shoes) (*shoe.Shoes, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.Shoes{}, err
	}

	return &shoes, nil
}

func AddShoesType(shoes shoe.ShoesType) (*shoe.ShoesType, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.ShoesType{}, err
	}

	return &shoes, nil
}
