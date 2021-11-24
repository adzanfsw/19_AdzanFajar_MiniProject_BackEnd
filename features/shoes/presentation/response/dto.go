package response

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

func FromShoes(core shoes.Shoes) Shoes {
	return Shoes{
		ID:          core.ID,
		Name:        core.Name,
		Price:       core.Price,
		BrandID:     core.BrandID,
		ShoesTypeID: core.ShoesTypeID,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
	}
}

func FromShoesSlice(core []shoes.Shoes) []Shoes {

	var shoeArray []Shoes

	for key := range core {
		shoeArray = append(shoeArray, FromShoes(core[key]))
	}

	return shoeArray
}
