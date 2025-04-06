package repository

import (
	"concerts/internal/models"
	"database/sql"
	"fmt"
)

type BookingRepository interface {
	GetBookings(partnerID int) ([]models.Booking, error)
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}

// GetBookings ดึงข้อมูลการจองของ partner
func (r *bookingRepository) GetBookings(partnerID int) ([]models.Booking, error) {
	var bookings []models.Booking
	rows, err := r.db.Query("SELECT id, concert_id, booking_at, booking_date, amount FROM bookings WHERE partner_id = $1", partnerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching bookings for partner %d: %w", partnerID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var booking models.Booking
		// สแกนค่าลงในฟิลด์ต่างๆ ของ Booking struct
		if err := rows.Scan(&booking.ID, &booking.ConcertID, &booking.BookingAt, &booking.BookingDate, &booking.Amount); err != nil {
			return nil, fmt.Errorf("error scanning booking data for partner %d: %w", partnerID, err)
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over bookings for partner %d: %w", partnerID, err)
	}

	return bookings, nil
}
