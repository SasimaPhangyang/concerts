package models

import "time"

type Booking struct {
	ID          int       `json:"id"`
	ConcertID   int       `json:"concert_id"`
	PartnerID   int       `json:"partner_id"`
	Tickets     int       `json:"tickets"`
	Amount      float64   `json:"amount"`
	BookingAt   time.Time `json:"booking_at"`
	BookingDate time.Time `json:"date"`
}
