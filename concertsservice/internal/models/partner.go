package models

import "time"

// PartnerReward คือข้อมูลรางวัลที่ได้รับจากพาร์ทเนอร์
type PartnerReward struct {
	RewardID  int       `json:"reward_id"`  // รหัสรางวัล
	Amount    float64   `json:"amount"`     // จำนวนเงินรางวัล
	PartnerID int       `json:"partner_id"` // รหัสพาร์ทเนอร์
	CreatedAt time.Time `json:"created_at"` // วันที่สร้างรางวัล
}
