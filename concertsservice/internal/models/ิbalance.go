package models

import "time"

// PartnerBalance คือข้อมูลยอดเงินของพาร์ทเนอร์
type PartnerBalance struct {
	PartnerID int       `json:"partner_id"` // รหัสพาร์ทเนอร์
	Balance   float64   `json:"balance"`    // ยอดเงินของพาร์ทเนอร์
	CreatedAt time.Time `json:"created_at"` // วันที่สร้าง
}
