package repository

import (
	"concerts/internal/models"
	"database/sql"
	"fmt"
)

type ReportRepository interface {
	GetSalesReport(product string) ([]models.SalesReport, error)
	GetSalesBySourceReport() ([]models.SalesBySource, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

// GetSalesReport ดึงข้อมูลยอดขายตาม product
func (r *reportRepository) GetSalesReport(product string) ([]models.SalesReport, error) {
	query := "SELECT product, amount, sales_date FROM sales_reports"
	var args []interface{}

	if product != "" {
		query += " WHERE product = $1"
		args = append(args, product)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching sales report: %w", err)
	}
	defer rows.Close()

	var reports []models.SalesReport
	for rows.Next() {
		var report models.SalesReport
		if err := rows.Scan(&report.Product, &report.Amount, &report.SalesDate); err != nil {
			return nil, fmt.Errorf("error scanning sales report: %w", err)
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over sales reports: %w", err)
	}

	return reports, nil
}

// GetSalesBySourceReport ดึงข้อมูลยอดขายแยกตาม source
func (r *reportRepository) GetSalesBySourceReport() ([]models.SalesBySource, error) {
	query := "SELECT source, total_sales FROM sales_by_source"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching sales by source report: %w", err)
	}
	defer rows.Close()

	var reports []models.SalesBySource
	for rows.Next() {
		var report models.SalesBySource
		if err := rows.Scan(&report.Source, &report.TotalSales); err != nil {
			return nil, fmt.Errorf("error scanning sales by source report: %w", err)
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over sales by source reports: %w", err)
	}

	return reports, nil
}
