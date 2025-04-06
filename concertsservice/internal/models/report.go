package models

// SalesReport ใช้สำหรับรายงานยอดขายโดยระบุวันที่และรหัสอีเวนต์
type SalesReport struct {
	Date        string  `json:"date"`         // วันที่
	EventID     string  `json:"event_id"`     // รหัสอีเวนต์
	SalesAmount float64 `json:"sales_amount"` // ยอดขาย
}

// SalesBySourceReport ใช้สำหรับรายงานยอดขายแยกตามแหล่งที่มา
type SalesBySourceReport struct {
	Source      string  `json:"source"`       // แหล่งที่มาของยอดขาย
	SalesAmount float64 `json:"sales_amount"` // ยอดขาย
}
