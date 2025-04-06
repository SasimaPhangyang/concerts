package models

import "time"

type Booking struct {
	ID        string    `json:"id"`
	ConcertID string    `json:"concert_id"` // เพิ่มฟิลด์ ConcertID
	Date      time.Time `json:"date"`
	Amount    float64   `json:"amount"`
}
