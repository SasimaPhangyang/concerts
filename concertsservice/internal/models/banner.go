package models

import "time"

type Banner struct {
	ID        int       `json:"id"`
	ImageURL  string    `json:"image_url"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
