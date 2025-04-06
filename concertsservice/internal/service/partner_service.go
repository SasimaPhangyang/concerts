package service

import (
	"concerts/internal/models"
	"concerts/internal/repository"
	"errors"
)

type PartnerService interface {
	GetPartnerBalance(partnerID string) (float64, error)
	SetAutoWithdraw(partnerID string, threshold float64, withdrawalMethod string) error
	RequestWithdrawal(partnerID string, amount float64) error
	GetPartnerRewards(partnerID string) ([]models.Reward, error)
	GetBookings(partnerID string) ([]models.Booking, error)
}

type partnerService struct {
	partnerRepo  repository.PartnerRepository
	bookingRepo  repository.BookingRepository
	withdrawRepo repository.WithdrawRepository
}

func NewPartnerService(partnerRepo repository.PartnerRepository, bookingRepo repository.BookingRepository, withdrawRepo repository.WithdrawRepository) PartnerService {
	return &partnerService{
		partnerRepo:  partnerRepo,
		bookingRepo:  bookingRepo,
		withdrawRepo: withdrawRepo,
	}
}

func (s *partnerService) GetPartnerBalance(partnerID string) (float64, error) {
	return s.partnerRepo.GetPartnerBalance(partnerID)
}

func (s *partnerService) GetPartnerRewards(partnerID string) ([]models.Reward, error) {
	rewards, err := s.partnerRepo.GetPartnerRewards(partnerID)
	if err != nil {
		return nil, err
	}
	return rewards, nil
}

func (s *partnerService) SetAutoWithdraw(partnerID string, threshold float64, withdrawalMethod string) error {
	if threshold <= 0 {
		return errors.New("threshold must be greater than zero")
	}
	if withdrawalMethod == "" {
		return errors.New("withdrawal method cannot be empty")
	}
	return s.partnerRepo.SetAutoWithdraw(partnerID, threshold, withdrawalMethod)
}

func (s *partnerService) RequestWithdrawal(partnerID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	// ตรวจสอบข้อมูลการถอน
	err := s.partnerRepo.RequestWithdrawal(partnerID, amount)
	if err != nil {
		return err
	}
	return nil
}

func (s *partnerService) GetBookings(partnerID string) ([]models.Booking, error) {
	bookings, err := s.bookingRepo.GetBookings(partnerID)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
