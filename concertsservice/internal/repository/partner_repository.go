package repository

import (
	"concerts/internal/models"
	"database/sql"
	"fmt"
)

type PartnerRepository interface {
	GetPartnerBalance(partnerID string) (float64, error)
	SetAutoWithdraw(partnerID string, threshold float64, withdrawalMethod string) error
	RequestWithdrawal(partnerID string, amount float64) error
	GetPartnerRewards(partnerID string) ([]models.Reward, error)
	GetBookings(partnerID string) ([]models.Booking, error) // เพิ่มที่นี่
}

type partnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) PartnerRepository {
	return &partnerRepository{db: db}
}

// GetPartnerBalance ดึงข้อมูล balance ของ partner จากฐานข้อมูล
func (r *partnerRepository) GetPartnerBalance(partnerID string) (float64, error) {
	var balance float64
	err := r.db.QueryRow("SELECT balance FROM partners WHERE id=$1", partnerID).Scan(&balance)

	if err != nil {
		return 0, fmt.Errorf("error fetching balance for partner %s: %w", partnerID, err)
	}
	return balance, nil
}

// SetAutoWithdraw ตั้งค่าการถอนอัตโนมัติของ partner
func (r *partnerRepository) SetAutoWithdraw(partnerID string, threshold float64, withdrawalMethod string) error {
	_, err := r.db.Exec("UPDATE partners SET auto_withdraw_threshold=$1, withdrawal_method=$2 WHERE partner_id=$3", threshold, withdrawalMethod, partnerID)
	if err != nil {
		return fmt.Errorf("error setting auto withdraw for partner %s: %w", partnerID, err)
	}
	return nil
}

// RequestWithdrawal ถอนเงินจาก balance ของ partner
func (r *partnerRepository) RequestWithdrawal(partnerID string, amount float64) error {
	// ตรวจสอบ balance ก่อนทำการถอน
	balance, err := r.GetPartnerBalance(partnerID)
	if err != nil {
		return err
	}

	if balance < amount {
		return fmt.Errorf("insufficient balance for partner %s: available balance %.2f, requested amount %.2f", partnerID, balance, amount)
	}

	// อัพเดต balance หลังจากถอน
	_, err = r.db.Exec("INSERT INTO withdraw_requests (partner_id, amount) VALUES ($1, $2)", partnerID, amount)
	if err != nil {
		return fmt.Errorf("error logging withdrawal for partner %s: %w", partnerID, err)
	}

	return nil
}

// GetPartnerRewards ดึงข้อมูลรางวัลของ partner
func (r *partnerRepository) GetPartnerRewards(partnerID string) ([]models.Reward, error) {
	var rewards []models.Reward
	rows, err := r.db.Query("SELECT reward_id, reward_name, amount FROM rewards WHERE partner_id = $1", partnerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching rewards for partner %s: %w", partnerID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var reward models.Reward
		if err := rows.Scan(&reward.ID, &reward.Name, &reward.Amount); err != nil {
			return nil, fmt.Errorf("error scanning reward data for partner %s: %w", partnerID, err)
		}
		rewards = append(rewards, reward)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rewards for partner %s: %w", partnerID, err)
	}

	if len(rewards) == 0 {
		return nil, fmt.Errorf("no rewards found for partner with ID %s", partnerID)
	}

	return rewards, nil
}

// GetBookings ดึงข้อมูลการจองของ partner
func (r *partnerRepository) GetBookings(partnerID string) ([]models.Booking, error) {
	var bookings []models.Booking
	rows, err := r.db.Query(`
    SELECT b.id, b.booking_at, b.amount
    FROM bookings b
    JOIN concerts c ON b.concert_id = c.id
    WHERE c.partner_id = $1
`, partnerID)

	if err != nil {
		return nil, fmt.Errorf("error fetching bookings for partner %s: %w", partnerID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.ID, &booking.Date, &booking.Amount); err != nil {
			return nil, fmt.Errorf("error scanning booking data for partner %s: %w", partnerID, err)
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over bookings for partner %s: %w", partnerID, err)
	}

	return bookings, nil
}
