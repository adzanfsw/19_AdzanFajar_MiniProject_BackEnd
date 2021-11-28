package shoes

import "time"

type Shoes struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	BrandID     int       `json:"brand_id"`
	ShoesTypeID int       `json:"shoes_type_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
