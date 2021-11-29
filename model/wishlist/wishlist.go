package wishlist

import "time"

type Wishlist struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ShoesID   int       `json:"shoes_id"`
	CreatedAt time.Time `json:"created_at"`
}
