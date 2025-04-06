package repository

import (
	"concerts/internal/models"
	"context"
	"database/sql"
	"errors"
)

type CommissionRepository interface {
	GetCommissions(ctx context.Context, partnerID int) ([]models.Commission, error)
}

type commissionRepository struct {
	db *sql.DB
}

func NewCommissionRepository(db *sql.DB) CommissionRepository {
	return &commissionRepository{db: db}
}

func (r *commissionRepository) GetCommissions(ctx context.Context, partnerID int) ([]models.Commission, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, partner_id, amount, date_from 
		FROM commissions 
		WHERE partner_id = $1 
		ORDER BY date_from`, partnerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commissions []models.Commission
	for rows.Next() {
		var c models.Commission
		if err := rows.Scan(&c.ID, &c.PartnerID, &c.Amount, &c.Date); err != nil {
			return nil, err
		}
		commissions = append(commissions, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(commissions) == 0 {
		return nil, errors.New("no commissions found for this partner")
	}

	return commissions, nil
}
