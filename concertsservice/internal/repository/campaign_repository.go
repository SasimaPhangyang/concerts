package repository

import (
	"database/sql"
	//"concerts/internal/models"
)

// CampaignRepository กำหนด interface สำหรับการเข้าร่วมแคมเปญ
type CampaignRepository interface {
	JoinCampaign(partnerID, campaignID string) error
}

type campaignRepository struct {
	db *sql.DB
}

// NewCampaignRepository ฟังก์ชันสำหรับสร้าง CampaignRepository
func NewCampaignRepository(db *sql.DB) CampaignRepository {
	return &campaignRepository{db: db}
}

// JoinCampaign ฟังก์ชันสำหรับเข้าร่วมแคมเปญ
func (r *campaignRepository) JoinCampaign(partnerID, campaignID string) error {
	// สร้าง SQL query สำหรับการเข้าร่วมแคมเปญ
	_, err := r.db.Exec(
		"INSERT INTO campaigns (partner_id, campaign_id) VALUES ($1, $2)",
		partnerID, campaignID,
	)

	// ถ้ามีข้อผิดพลาดในการแทรกข้อมูล ให้คืนค่า error
	if err != nil {
		return err
	}

	// ถ้าทุกอย่างเรียบร้อย ไม่มีข้อผิดพลาด
	return nil
}
