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
