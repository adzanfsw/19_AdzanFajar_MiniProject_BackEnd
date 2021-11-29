package database

import (
	"justrun/config"
	"justrun/model/users"
)

func AddUsers(user users.Users) (*users.Users, error) {

	if err := config.DB.Save(&user).Error; err != nil {
		return &users.Users{}, err
	}

	return &user, nil
}

func GetUsers() (*[]users.Users, error) {

	var user []users.Users

	if err := config.DB.Find(&user).Error; err != nil {
		return &[]users.Users{}, err
	}

	return &user, nil
}

func UserByID(id int) (*users.Users, error) {
	var user users.Users

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &users.Users{}, err
	}

	return &user, nil
}

func UpdateUsers(id int, u users.Users) (*users.Users, error) {

	if err := config.DB.Where("id = ?", id).Updates(&u).Error; err != nil {
		return &users.Users{}, err
	}

	return &u, nil
}

func DeleteUsers(id int) (*users.Users, error) {

	var user users.Users

	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return &users.Users{}, err
	}

	return &user, nil
}
