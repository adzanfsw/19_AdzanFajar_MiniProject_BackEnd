package config

import (
	"justrun/model/review"
	"justrun/model/shoes"
	"justrun/model/users"
	"justrun/model/wishlist"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbname string) {

	var err error

	db, err := gorm.Open(mysql.Open("root:@/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

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

		&wishlist.Wishlist{},

		&review.Review{},
	)
}
