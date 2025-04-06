package models

import "time"

// AutoWithdrawRequest คือการร้องขอการตั้งค่าการถอนอัตโนมัติ
type AutoWithdrawRequest struct {
	Enabled bool `json:"enabled"` // ใช้ bool เพื่อระบุสถานะการเปิด/ปิดการถอนอัตโนมัติ
}

// PartnerReward คือข้อมูลรางวัลที่ได้รับจากพาร์ทเนอร์
type PartnerReward struct {
	RewardID  int       `json:"reward_id"`  // รหัสรางวัล
	Amount    float64   `json:"amount"`     // จำนวนเงินรางวัล
	PartnerID int       `json:"partner_id"` // รหัสพาร์ทเนอร์
	CreatedAt time.Time `json:"created_at"` // วันที่สร้างรางวัล
}
