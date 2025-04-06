package service

import (
	"concerts/internal/models"
	"concerts/internal/repository"
	"fmt"
)

type ReportService interface {
	GetSalesReport(product string) ([]models.SalesReport, error)
	GetSalesBySourceReport() ([]models.SalesBySource, error)
}

type reportService struct {
	repo repository.ReportRepository
}

// Inject Repository ผ่าน Constructor ตามหลัก DIP
func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

// GetSalesReport ดึงข้อมูลยอดขายตาม product
func (s *reportService) GetSalesReport(product string) ([]models.SalesReport, error) {
	// ตรวจสอบว่า product ต้องมีค่า
	if product == "" {
		return nil, fmt.Errorf("product must be provided")
	}
	return s.repo.GetSalesReport(product)
}

// GetSalesBySourceReport ดึงข้อมูลยอดขายแยกตาม source
func (s *reportService) GetSalesBySourceReport() ([]models.SalesBySource, error) {
	return s.repo.GetSalesBySourceReport()
}
