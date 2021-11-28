package shoes

import "time"

type ShoesDescription struct {
	ID          int       `json:"id"`
	ShoesID     int       `json:"shoes_id"`
	Description string    `json:"desc"`
	Color       string    `json:"color"`
	PurchaseURL string    `json:"purchase_url"`
	CreatedAt   time.Time `json:"created_at"`
}
