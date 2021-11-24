package data

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

func (a *Shoes) toCore() shoes.Shoes {

	return shoes.Shoes{
		ID:          int(a.ID),
		Name:        a.Name,
		Price:       a.Price,
		BrandID:     a.BrandID,
		ShoesTypeID: a.ShoesTypeID,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func toShoesList(resp []Shoes) []shoes.Shoes {

	a := []shoes.Shoes{}

	for key := range resp {
		a = append(a, resp[key].toCore())
	}

	return a
}

func fromShoes(core shoes.Shoes) Shoes {
	return Shoes{}
}

// func toRecordShoes(data shoes.Shoes) Shoes {

// 	convertedRequirement := []Requirement{}

// 	for _, req := range data.Requirements {
// 		convertedRequirement = append(convertedRequirement, toRecordRequirement(req))
// 	}

// 	return Shoes{
// 		Name        data.Name,
// 		Price       data.Price,
// 		BrandID     data.BrandID,
// 		ShoesTypeID data.ShoesTypeID,
// }
