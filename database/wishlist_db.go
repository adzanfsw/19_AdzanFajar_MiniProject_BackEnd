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
