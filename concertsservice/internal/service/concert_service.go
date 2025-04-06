package service

import (
	"concerts/internal/models"
	"concerts/internal/repository"
)

type ConcertService interface {
	GetAllConcerts() ([]models.Concert, error)
	GetConcertByID(id int) (*models.Concert, error)
	SearchConcerts(query string) ([]models.Concert, error) // เพิ่ม method นี้
}

type concertService struct {
	repo repository.ConcertRepository
}

func NewConcertService(repo repository.ConcertRepository) ConcertService {
	return &concertService{repo: repo}
}

func (s *concertService) GetAllConcerts() ([]models.Concert, error) {
	return s.repo.GetAllConcerts()
}

func (s *concertService) GetConcertByID(id int) (*models.Concert, error) {
	return s.repo.GetConcertByID(id)
}

// เพิ่มการค้นหาคอนเสิร์ต
func (s *concertService) SearchConcerts(query string) ([]models.Concert, error) {
	return s.repo.SearchConcerts(query)
}
