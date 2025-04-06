package repository

import (
	"concerts/internal/models"
	"database/sql"
	"fmt"
)

type ReportRepository interface {
	GetSalesReport(date, eventID string) ([]models.SalesReport, error)
	GetSalesBySourceReport(month string) ([]models.SalesBySourceReport, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

// GetSalesReport ดึงข้อมูลรายงานการขายตามวันที่หรือช่วงวันที่ และอีเวนต์
func (r *reportRepository) GetSalesReport(date, eventID string) ([]models.SalesReport, error) {
	var query string
	var args []interface{}

	// สร้างคำสั่ง SQL ตามพารามิเตอร์ที่ได้รับ
	if date != "" {
		query = "SELECT date, event_id, sales_amount FROM sales_reports WHERE date = $1"
		args = append(args, date)
	}
	if eventID != "" {
		if query != "" {
			query += " AND event_id = $2"
		} else {
			query = "SELECT date, event_id, sales_amount FROM sales_reports WHERE event_id = $1"
		}
		args = append(args, eventID)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching sales report: %w", err)
	}
	defer rows.Close()

	var reports []models.SalesReport
	for rows.Next() {
		var report models.SalesReport
		if err := rows.Scan(&report.Date, &report.EventID, &report.SalesAmount); err != nil {
			return nil, fmt.Errorf("error scanning sales report: %w", err)
		}
		reports = append(reports, report)
	}

	// ตรวจสอบข้อผิดพลาดจากการวนลูป
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over sales reports: %w", err)
	}

	return reports, nil
}

// GetSalesBySourceReport ดึงข้อมูลรายงานการขายตามแหล่งที่มาในเดือนที่กำหนด
func (r *reportRepository) GetSalesBySourceReport(month string) ([]models.SalesBySourceReport, error) {
	// ตรวจสอบค่าของ month ก่อน
	if month == "" {
		return nil, fmt.Errorf("month cannot be empty")
	}

	// สร้างคำสั่ง SQL เพื่อค้นหาข้อมูลตามเดือนที่กำหนด
	query := "SELECT source, sales_amount FROM sales_by_source WHERE TO_CHAR(date, 'YYYY-MM') = $1"
	rows, err := r.db.Query(query, month)
	if err != nil {
		return nil, fmt.Errorf("error fetching sales by source report: %w", err)
	}
	defer rows.Close()

	var reports []models.SalesBySourceReport
	for rows.Next() {
		var report models.SalesBySourceReport
		if err := rows.Scan(&report.Source, &report.SalesAmount); err != nil {
			return nil, fmt.Errorf("error scanning sales by source report: %w", err)
		}
		reports = append(reports, report)
	}

	// ตรวจสอบข้อผิดพลาดจากการวนลูป
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over sales by source reports: %w", err)
	}

	return reports, nil
}
