package database

import (
	"justrun/config"
	shoe "justrun/model/shoes"
)

func AddShoesBrand(shoes shoe.ShoesBrand) (*shoe.ShoesBrand, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.ShoesBrand{}, err
	}

	return &shoes, nil
}
