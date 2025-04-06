package service

import (
	"concerts/internal/models"
	"concerts/internal/repository"
)

type BannerService interface {
	GetBanners() ([]models.Banner, error)
}

type bannerService struct {
	repo repository.BannerRepository
}

// Inject Repository ผ่าน Constructor ตามหลัก DIP
func NewBannerService(repo repository.BannerRepository) BannerService {
	return &bannerService{repo: repo}
}

// แก้ไขให้เรียกใช้ GetAllBanners จาก BannerRepository
func (s *bannerService) GetBanners() ([]models.Banner, error) {
	// เรียกใช้ฟังก์ชัน GetAllBanners จาก BannerRepository
	return s.repo.GetAllBanners()
}
