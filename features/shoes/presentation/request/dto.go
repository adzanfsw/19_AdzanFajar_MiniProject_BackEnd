package request

import (
	"justrun/features/shoes"
	"time"
)

type Shoes struct {
	ID          int
	Name        string
	Price       int
	BrandID     int `json:"brand_id"`
	ShoesTypeID int `json:"shoes_type_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToShoes(req Shoes) shoes.Shoes {
	return shoes.Shoes{
		ID:          req.ID,
		Name:        req.Name,
		Price:       req.Price,
		BrandID:     req.BrandID,
		ShoesTypeID: req.ShoesTypeID,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
}
