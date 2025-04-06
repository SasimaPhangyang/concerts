package repository

import (
	"database/sql"
	"fmt"
)

type WithdrawRepository interface {
	// สร้าง method สำหรับการจัดการการถอน
	RequestWithdrawal(partnerID string, amount float64) error
}

type withdrawRepository struct {
	db *sql.DB
}

func NewWithdrawRepository(db *sql.DB) WithdrawRepository {
	return &withdrawRepository{db: db}
}

// RequestWithdrawal ทำการถอนเงิน
func (r *withdrawRepository) RequestWithdrawal(partnerID string, amount float64) error {
	// สมมติว่ามี SQL สำหรับการบันทึกการถอน
	_, err := r.db.Exec("INSERT INTO withdrawals (partner_id, amount) VALUES ($1, $2)", partnerID, amount)
	if err != nil {
		return fmt.Errorf("error requesting withdrawal for partner %s: %w", partnerID, err)
	}
	return nil
}
