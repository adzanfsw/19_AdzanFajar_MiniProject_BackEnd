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

type Bussiness interface {
	CreateData(data Shoes) (resp Shoes, err error)
	GetAllData(search string) (resp []Shoes)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Shoes) (resp Shoes, err error)
	SelectData(title string) (resp []Shoes)
}
