package review

import "time"

type Review struct {
	ID        int       `json:"id"`
	Rating    int       `json:"rating"`
	Review    string    `json:"review"`
	UserID    int       `json:"user_id"`
	ShoesID   int       `json:"shoes_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
