package service

import (
	"concerts/internal/models"
	"concerts/internal/repository"
	"fmt"
)

type ReportService interface {
	GetSalesReport(date, eventID string) ([]models.SalesReport, error)
	GetSalesBySourceReport(month string) ([]models.SalesBySourceReport, error)
}

type reportService struct {
	repo repository.ReportRepository
}

// Inject Repository ผ่าน Constructor ตามหลัก DIP
func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

// GetSalesReport ดึงข้อมูลรายงานการขายตามวันที่หรือช่วงวันที่
func (s *reportService) GetSalesReport(date, eventID string) ([]models.SalesReport, error) {
	// ตรวจสอบว่าอย่างน้อยต้องมี date หรือ event_id ที่ถูกส่งมา
	if date == "" && eventID == "" {
		return nil, fmt.Errorf("either date or event_id must be provided")
	}
	// เรียกใช้ repository เพื่อดึงข้อมูล
	return s.repo.GetSalesReport(date, eventID)
}

// GetSalesBySourceReport ดึงข้อมูลรายงานการขายตามแหล่งที่มา
func (s *reportService) GetSalesBySourceReport(month string) ([]models.SalesBySourceReport, error) {
	// ตรวจสอบว่าเดือนต้องมีค่า
	if month == "" {
		return nil, fmt.Errorf("month must be provided")
	}
	// เรียกใช้ repository เพื่อดึงข้อมูล
	return s.repo.GetSalesBySourceReport(month)
}
