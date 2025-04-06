package service

import (
	"concerts/internal/repository"
	"errors"
)

type CampaignService interface {
	JoinCampaign(partnerID, campaignID string) error
}

type campaignService struct {
	repo repository.CampaignRepository
}

// Inject Repository ผ่าน Constructor ตามหลัก DIP
func NewCampaignService(repo repository.CampaignRepository) CampaignService {
	return &campaignService{repo: repo}
}

// JoinCampaign ใช้ในการเข้าร่วมแคมเปญ
func (s *campaignService) JoinCampaign(partnerID, campaignID string) error {
	if partnerID == "" || campaignID == "" {
		return errors.New("partner_id and campaign_id are required")
	}

	// เรียกใช้ repository เพื่อลงทะเบียนเข้าร่วมแคมเปญ
	err := s.repo.JoinCampaign(partnerID, campaignID)
	if err != nil {
		// ถ้ามีข้อผิดพลาดจาก repository ให้คืนค่า error
		return err
	}

	// ถ้าทุกอย่างเรียบร้อย ให้คืนค่า nil (ไม่มีข้อผิดพลาด)
	return nil
}
