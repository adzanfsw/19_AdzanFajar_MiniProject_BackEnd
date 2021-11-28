package shoes

import "time"

type ShoesDesc struct {
	ShoesID     int       `json:"shoes_id"`
	Description string    `json:"desc"`
	Color       string    `json:"color"`
	PurchaseURL string    `json:"purchase_url"`
	CreatedAt   time.Time `json:"created_at"`
}
