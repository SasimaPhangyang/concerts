package repository

import (
	"concerts/internal/models"
	"database/sql"
	"fmt"
)

type PartnerRepository interface {
	GetPartnerBalance(partnerID int) (float64, error)
	GetAutoWithdrawSetting(partnerID int) (models.AutoWithdraw, error)
	SetAutoWithdraw(partnerID int, enabled bool) error
	RequestWithdrawal(partnerID int, amount float64) error
	GetPartnerRewards(partnerID int) ([]models.Reward, error)
	GetBookings(partnerID int) ([]models.Booking, error)
}

type partnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) PartnerRepository {
	return &partnerRepository{db: db}
}

// GetPartnerBalance ดึงข้อมูล balance ของ partner จากฐานข้อมูล
func (r *partnerRepository) GetPartnerBalance(partnerID int) (float64, error) {
	var balance float64
	err := r.db.QueryRow("SELECT balance FROM partners WHERE id=$1", partnerID).Scan(&balance)

	if err != nil {
		return 0, fmt.Errorf("error fetching balance for partner %d: %w", partnerID, err)
	}
	return balance, nil
}

// GetAutoWithdrawSetting
func (r *partnerRepository) GetAutoWithdrawSetting(partnerID int) (models.AutoWithdraw, error) {
	var autoWithdraw models.AutoWithdraw
	query := "SELECT partner_id, enabled FROM auto_withdraw WHERE partner_id = $1"
	err := r.db.QueryRow(query, partnerID).Scan(&autoWithdraw.PartnerID, &autoWithdraw.Enabled)
	if err != nil {
		if err == sql.ErrNoRows {
			return autoWithdraw, fmt.Errorf("auto withdraw setting not found for partner %d", partnerID)
		}
		return autoWithdraw, fmt.Errorf("error fetching auto withdraw setting for partner %d: %w", partnerID, err)
	}
	return autoWithdraw, nil
}

// SetAutoWithdraw
func (r *partnerRepository) SetAutoWithdraw(partnerID int, enabled bool) error {
	_, err := r.db.Exec("INSERT INTO auto_withdraw (partner_id, enabled) VALUES ($1, $2) ON CONFLICT (partner_id) DO UPDATE SET enabled = $2", partnerID, enabled)
	if err != nil {
		return fmt.Errorf("error setting auto withdraw for partner %d: %w", partnerID, err)
	}
	return nil
}

// RequestWithdrawal ขอถอนเงิน
func (r *partnerRepository) RequestWithdrawal(partnerID int, amount float64) error {
	// ตรวจสอบการตั้งค่าการถอนอัตโนมัติ
	autoWithdraw, err := r.GetAutoWithdrawSetting(partnerID)
	if err != nil {
		return fmt.Errorf("error fetching auto withdraw setting: %w", err)
	}

	// ถ้าการถอนอัตโนมัติไม่ได้เปิดใช้งาน
	if !autoWithdraw.Enabled {
		return fmt.Errorf("auto withdraw is not enabled for partner %d", partnerID)
	}

	_, err = r.db.Exec("INSERT INTO withdraw_requests (partner_id, amount) VALUES ($1, $2)", partnerID, amount)
	if err != nil {
		return fmt.Errorf("error requesting withdrawal for partner %d: %w", partnerID, err)
	}
	return nil
}

// GetPartnerRewards ดึงข้อมูลรางวัลของ partner
func (r *partnerRepository) GetPartnerRewards(partnerID int) ([]models.Reward, error) {
	var rewards []models.Reward
	rows, err := r.db.Query("SELECT reward_id, amount, partner_id, created_at FROM partner_rewards WHERE partner_id = $1", partnerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching rewards for partner %d: %w", partnerID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var reward models.Reward
		if err := rows.Scan(&reward.ID, &reward.Amount, &reward.PartnerID, &reward.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning reward data for partner %d: %w", partnerID, err)
		}
		rewards = append(rewards, reward)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rewards for partner %d: %w", partnerID, err)
	}

	return rewards, nil
}

// GetBookings ดึงข้อมูลการจองของ partner
func (r *partnerRepository) GetBookings(partnerID int) ([]models.Booking, error) {
	var bookings []models.Booking
	rows, err := r.db.Query(`
		SELECT id, concert_id, partner_id, tickets, amount, booking_at, date 
		FROM bookings 
		WHERE partner_id = $1`, partnerID)

	if err != nil {
		return nil, fmt.Errorf("error fetching bookings for partner %d: %w", partnerID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(
			&booking.ID,
			&booking.ConcertID,
			&booking.PartnerID,
			&booking.Tickets,
			&booking.Amount,
			&booking.BookingAt,
			&booking.BookingDate,
		); err != nil {
			return nil, fmt.Errorf("error scanning booking data for partner %d: %w", partnerID, err)
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over bookings for partner %d: %w", partnerID, err)
	}

	return bookings, nil
}
