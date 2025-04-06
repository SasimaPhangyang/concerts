package models

import "time"

type Reward struct {
	ID        int       `json:"reward_id"`
	Amount    float64   `json:"amount"`
	PartnerID int       `json:"partner_id"`
	CreatedAt time.Time `json:"created_at"`
}
