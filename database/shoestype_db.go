package database

import (
	"justrun/config"
	shoe "justrun/model/shoes"
)

func AddShoesType(shoes shoe.ShoesType) (*shoe.ShoesType, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.ShoesType{}, err
	}

	return &shoes, nil
}
