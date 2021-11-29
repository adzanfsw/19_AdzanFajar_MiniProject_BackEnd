package config

import (
	"justrun/model/shoes"
	"justrun/model/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error

	db, err := gorm.Open(mysql.Open("root:@/justrun?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}

func AutoMigrate() {

	DB.AutoMigrate(
		&shoes.Shoes{},
		&shoes.ShoesType{},
		&shoes.ShoesBrand{},
		&shoes.ShoesDescription{},

		&users.Users{},
	)
}
