package shoes

import "time"

type ShoesType struct {
	ID        int
	Type      string `json:"type"`
	CreatedAt time.Time
}
