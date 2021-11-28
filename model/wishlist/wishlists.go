package wishlist

import "time"

type Wishlist struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	ShoesID   string    `json:"shoes_id"`
	CreatedAt time.Time `json:"created_at"`
}
