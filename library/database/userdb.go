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
