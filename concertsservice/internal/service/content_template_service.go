package service

import (
	"concerts/internal/models"
	"concerts/internal/repository"
)

type ContentTemplateService interface {
	GetContentTemplates() ([]models.ContentTemplate, error)
}

type contentTemplateService struct {
	repo repository.ContentTemplateRepository
}

// Inject Repository ผ่าน Constructor ตามหลัก DIP
func NewContentTemplateService(repo repository.ContentTemplateRepository) ContentTemplateService {
	return &contentTemplateService{repo: repo}
}

func (s *contentTemplateService) GetContentTemplates() ([]models.ContentTemplate, error) {
	// เรียกใช้เมธอด GetAllTemplates ที่ถูกสร้างใน repository
	return s.repo.GetAllTemplates()
}
