package repository

import (
	"concerts/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type BookingRepository interface {
	GetBookings(ctx context.Context, partnerID int) ([]models.Booking, error)
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}

// GetBookings ดึงข้อมูลการจองของ partner
func (r *bookingRepository) GetBookings(ctx context.Context, partnerID int) ([]models.Booking, error) {
	var bookings []models.Booking

	query := `SELECT id, concert_id, partner_id, tickets, amount, booking_at, date
	          FROM bookings WHERE partner_id = ?`

	rows, err := r.db.QueryContext(ctx, query, partnerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching bookings for partner %d: %w", partnerID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.ID, &booking.ConcertID, &booking.PartnerID, &booking.Tickets, &booking.Amount, &booking.BookingAt, &booking.BookingDate); err != nil {
			return nil, fmt.Errorf("error scanning booking row for partner %d: %w", partnerID, err)
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over bookings for partner %d: %w", partnerID, err)
	}

	return bookings, nil
}
