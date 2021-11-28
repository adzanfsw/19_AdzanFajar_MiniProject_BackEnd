package shoes

import "time"

type ShoesType struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
