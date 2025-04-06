package repository

import (
	"concerts/internal/models"
	"database/sql"
	"fmt"
)

type BannerRepository interface {
	GetAllBanners() ([]models.Banner, error)
}

type bannerRepository struct {
	db *sql.DB
}

func NewBannerRepository(db *sql.DB) BannerRepository {
	return &bannerRepository{db: db}
}

func (r *bannerRepository) GetAllBanners() ([]models.Banner, error) {
	// Query เพื่อดึงข้อมูลแบนเนอร์จากฐานข้อมูล
	rows, err := r.db.Query("SELECT id, image_url, link, created_at, updated_at FROM banners ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch banners: %v", err)
	}
	defer rows.Close()

	var banners []models.Banner
	// Loop เพื่อดึงข้อมูลจากแต่ละแถว
	for rows.Next() {
		var b models.Banner
		if err := rows.Scan(&b.ID, &b.ImageURL, &b.Link, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan banner: %v", err)
		}
		banners = append(banners, b)
	}
	return banners, rows.Err()
}
