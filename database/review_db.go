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
