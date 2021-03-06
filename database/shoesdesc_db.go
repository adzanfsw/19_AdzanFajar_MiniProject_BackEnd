package database

import (
	"justrun/config"
	"justrun/model/shoes"
	shoe "justrun/model/shoes"
)

func AddShoesDescription(shoes shoe.ShoesDescription) (*shoe.ShoesDescription, error) {

	if err := config.DB.Save(&shoes).Error; err != nil {
		return &shoe.ShoesDescription{}, err
	}

	return &shoes, nil
}

func GetShoesDesc() (*[]shoes.ShoesDescription, error) {

	var desc []shoe.ShoesDescription

	if err := config.DB.Find(&desc).Error; err != nil {
		return &[]shoes.ShoesDescription{}, err
	}

	return &desc, nil
}

func DescByShoesID(id int) (*[]shoes.ShoesDescription, error) {

	var des []shoes.ShoesDescription

	if err := config.DB.Where("shoes_id = ?", id).Find(&des).Error; err != nil {
		return &[]shoes.ShoesDescription{}, err
	}

	return &des, nil
}
