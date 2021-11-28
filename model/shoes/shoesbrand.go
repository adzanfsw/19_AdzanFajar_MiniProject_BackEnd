package shoes

import "time"

type ShoesBrand struct {
	ID        int       `json:"id"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"created_at"`
}
