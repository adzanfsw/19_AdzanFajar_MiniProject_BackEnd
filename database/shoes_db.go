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

func UpdateShoes(id int, s shoe.Shoes) (*shoe.Shoes, error) {

	if err := config.DB.Where("id = ?", id).Updates(&s).Error; err != nil {
		return &shoe.Shoes{}, err
	}

	return &s, nil
}

// func DeleteShoes(id int) (*shoe.Shoes, error) {

// 	var sepatu shoe.Shoes

// 	if err := config.DB.Where("id = ?", id).Delete(&sepatu).Error; err != nil {
// 		return &shoe.Shoes{}, err
// 	}

// 	return &sepatu, nil
// }
