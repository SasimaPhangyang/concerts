package models

import "time"

type Commission struct {
	ID        int       `json:"id"`
	PartnerID string    `json:"partner_id"`
	Amount    float64   `json:"amount"`
	Date      time.Time `json:"date"`
}
