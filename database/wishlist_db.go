package database

import (
	"justrun/config"
	"justrun/model/wishlist"
)

func AddWishlist(wish wishlist.Wishlist) (*wishlist.Wishlist, error) {

	if err := config.DB.Save(&wish).Error; err != nil {
		return &wishlist.Wishlist{}, err
	}

	return &wish, nil
}

func GetWishlists() (*[]wishlist.Wishlist, error) {

	var wi []wishlist.Wishlist

	if err := config.DB.Find(&wi).Error; err != nil {
		return &[]wishlist.Wishlist{}, err
	}

	return &wi, nil
}

func WishByUserID(id int) (*[]wishlist.Wishlist, error) {

	var wis []wishlist.Wishlist

	if err := config.DB.Where("user_id = ?", id).Find(&wis).Error; err != nil {
		return &[]wishlist.Wishlist{}, err
	}

	return &wis, nil
}

func WishByShoesID(id int) (*[]wishlist.Wishlist, error) {

	var wis []wishlist.Wishlist

	if err := config.DB.Where("shoes_id = ?", id).Find(&wis).Error; err != nil {
		return &[]wishlist.Wishlist{}, err
	}

	return &wis, nil
}
