package database

import (
	"justrun/config"
	"justrun/model/shoes"
	shoe "justrun/model/shoes"
)

func AddShoesType(shoes shoe.ShoesType) (*shoe.ShoesType, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.ShoesType{}, err
	}

	return &shoes, nil
}

func GetShoesType() (*[]shoes.ShoesType, error) {

	var tipe []shoes.ShoesType

	if err := config.DB.Find(&tipe).Error; err != nil {
		return &[]shoes.ShoesType{}, err
	}

	return &tipe, nil
}

func ShoesTypeID(id int) (*shoes.ShoesType, error) {

	var tip shoes.ShoesType

	if err := config.DB.Where("id = ?", id).First(&tip).Error; err != nil {
		return &shoes.ShoesType{}, err
	}

	return &tip, nil
}
