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

func GetShoes() (*[]shoe.Shoes, error) {

	var shoes []shoe.Shoes

	if err := config.DB.Find(&shoes).Error; err != nil {
		return &[]shoe.Shoes{}, err
	}

	return &shoes, nil
}
