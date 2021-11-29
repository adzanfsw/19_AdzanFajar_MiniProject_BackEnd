package database

import (
	"justrun/config"
	"justrun/model/shoes"
	shoe "justrun/model/shoes"
)

func AddShoesBrand(shoes shoe.ShoesBrand) (*shoe.ShoesBrand, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.ShoesBrand{}, err
	}

	return &shoes, nil
}

func GetShoesBrand() (*[]shoes.ShoesBrand, error) {

	var merk []shoe.ShoesBrand

	if err := config.DB.Find(&merk).Error; err != nil {
		return &[]shoes.ShoesBrand{}, err
	}

	return &merk, nil
}
