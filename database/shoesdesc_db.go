package database

import (
	"justrun/config"
	shoe "justrun/model/shoes"
)

func AddShoesDescription(shoes shoe.ShoesDescription) (*shoe.ShoesDescription, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.ShoesDescription{}, err
	}

	return &shoes, nil
}
