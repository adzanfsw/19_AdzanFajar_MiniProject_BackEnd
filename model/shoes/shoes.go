package shoes

import "time"

type Shoes struct {
	ID          int
	Name        string
	Price       int
	BrandID     int `json:"brand_id"`
	ShoesTypeID int `json:"shoes_type_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
