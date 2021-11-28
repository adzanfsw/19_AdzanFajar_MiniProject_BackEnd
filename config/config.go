package config

import (
	"justrun/model/shoes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error

	db, err := gorm.Open(mysql.Open("root:@/justrun"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}

func AutoMigrate() {

	DB.AutoMigrate(
		&shoes.Shoes{},
		&shoes.ShoesType{},
	)
}
