package models

import "time"

type WithdrawRequest struct {
	ID        int       `json:"id"`
	PartnerID int       `json:"partner_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
