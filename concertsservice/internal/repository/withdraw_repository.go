package repository

import (
	"concerts/internal/models"
	"database/sql"
	"fmt"
)

type WithdrawRepository interface {
	CreateWithdrawRequest(partnerID int, amount float64) error
	GetWithdrawRequests(partnerID int) ([]models.WithdrawRequest, error)
}

type withdrawRepository struct {
	db *sql.DB
}

func NewWithdrawRepository(db *sql.DB) WithdrawRepository {
	return &withdrawRepository{db: db}
}

// CreateWithdrawRequest ทำการเพิ่มการถอนเงินใหม่
func (r *withdrawRepository) CreateWithdrawRequest(partnerID int, amount float64) error {
	_, err := r.db.Exec(`
		INSERT INTO withdraw_requests (partner_id, amount) 
		VALUES ($1, $2)`, partnerID, amount)
	if err != nil {
		return fmt.Errorf("error creating withdraw request for partner %d: %w", partnerID, err)
	}
	return nil
}

// GetWithdrawRequests ดึงรายการการถอนทั้งหมดสำหรับ partner
func (r *withdrawRepository) GetWithdrawRequests(partnerID int) ([]models.WithdrawRequest, error) {
	rows, err := r.db.Query(`
		SELECT id, partner_id, amount, created_at 
		FROM withdraw_requests 
		WHERE partner_id = $1`, partnerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching withdraw requests for partner %d: %w", partnerID, err)
	}
	defer rows.Close()

	var requests []models.WithdrawRequest
	for rows.Next() {
		var request models.WithdrawRequest
		if err := rows.Scan(&request.ID, &request.PartnerID, &request.Amount, &request.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning withdraw request data for partner %d: %w", partnerID, err)
		}
		requests = append(requests, request)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}
