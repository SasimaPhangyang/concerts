package models

// SalesReport
type SalesReport struct {
	Product   string  `json:"product"`    // ชื่อสินค้า
	Amount    float64 `json:"amount"`     // ยอดขาย
	SalesDate string  `json:"sales_date"` // วันที่ขาย (timestamp)
}

// SalesBySource
type SalesBySource struct {
	Source     string  `json:"source"`      // แหล่งที่มา
	TotalSales float64 `json:"total_sales"` // ยอดขายรวม
}
