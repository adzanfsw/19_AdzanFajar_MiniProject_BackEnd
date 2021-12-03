package database

import (
	"justrun/config"
	"justrun/model/review"
)

func AddReview(rev review.Review) (*review.Review, error) {

	if err := config.DB.Save(&rev).Error; err != nil {
		return &review.Review{}, err
	}

	return &rev, nil
}

func GetReviews() (*[]review.Review, error) {

	var revi []review.Review

	if err := config.DB.Find(&revi).Error; err != nil {
		return &[]review.Review{}, err
	}

	return &revi, nil
}

func UpdateReview(id int, re review.Review) (*review.Review, error) {

	if err := config.DB.Where("id = ?", id).Updates(&re).Error; err != nil {
		return &review.Review{}, err
	}

	return &re, nil
}

func ReviewByID(id int) (*[]review.Review, error) {

	var revie []review.Review

	if err := config.DB.Where("id = ?", id).Find(&revie).Error; err != nil {
		return &[]review.Review{}, err
	}

	return &revie, nil
}

func ReviewByShoesID(id int) (*[]review.Review, error) {

	var revie []review.Review

	if err := config.DB.Where("shoes_id = ?", id).Find(&revie).Error; err != nil {
		return &[]review.Review{}, err
	}

	return &revie, nil
}

func ReviewByRating(id int) (*[]review.Review, error) {

	var revie []review.Review

	if err := config.DB.Where("rating = ?", id).Find(&revie).Error; err != nil {
		return &[]review.Review{}, err
	}

	return &revie, nil
}

func ReviewByUserID(id int) (*[]review.Review, error) {

	var revie []review.Review

	if err := config.DB.Where("user_id = ?", id).Find(&revie).Error; err != nil {
		return &[]review.Review{}, err
	}

	return &revie, nil
}
