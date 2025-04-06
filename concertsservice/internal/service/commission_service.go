package service

import (
	"context"

	"concerts/internal/models"
	"concerts/internal/repository"
)

type CommissionService interface {
	GetCommissions(ctx context.Context, partnerID int) ([]models.Commission, error)
}

type commissionService struct {
	repo repository.CommissionRepository
}

// Inject Repository ผ่าน Constructor ตามหลัก DIP
func NewCommissionService(repo repository.CommissionRepository) CommissionService {
	return &commissionService{repo: repo}
}

func (s *commissionService) GetCommissions(ctx context.Context, partnerID int) ([]models.Commission, error) {
	return s.repo.GetCommissions(ctx, partnerID)
}
