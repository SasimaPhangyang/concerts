package service

import (
	"concerts/internal/models"
	"concerts/internal/repository"
	"context"
	"errors"
	"fmt"
)

type PartnerService interface {
	GetPartnerBalance(partnerID int) (float64, error)
	GetAutoWithdrawSetting(partnerID int) (models.AutoWithdraw, error)
	SetAutoWithdraw(partnerID int, enabled bool) error
	RequestWithdrawal(partnerID int, amount float64) error
	GetPartnerRewards(partnerID int) ([]models.Reward, error)
	GetBookings(partnerID int) ([]models.Booking, error)
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

func (s *partnerService) GetPartnerBalance(partnerID int) (float64, error) {
	return s.partnerRepo.GetPartnerBalance(partnerID)
}

func (s *partnerService) GetAutoWithdrawSetting(partnerID int) (models.AutoWithdraw, error) {
	autoWithdraw, err := s.partnerRepo.GetAutoWithdrawSetting(partnerID)
	if err != nil {
		return autoWithdraw, err
	}
	return autoWithdraw, nil
}

func (s *partnerService) SetAutoWithdraw(partnerID int, enabled bool) error {
	err := s.partnerRepo.SetAutoWithdraw(partnerID, enabled)
	if err != nil {
		return err
	}
	return nil
}

func (s *partnerService) RequestWithdrawal(partnerID int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	autoWithdraw, err := s.partnerRepo.GetAutoWithdrawSetting(partnerID)
	if err != nil {
		return fmt.Errorf("error fetching auto withdraw setting: %w", err)
	}

	// ถ้าการถอนอัตโนมัติไม่ได้เปิดใช้งาน
	if !autoWithdraw.Enabled {
		return fmt.Errorf("auto withdraw is not enabled for partner %d", partnerID)
	}

	// ทำการขอถอนเงิน
	err = s.partnerRepo.RequestWithdrawal(partnerID, amount)
	if err != nil {
		return fmt.Errorf("error requesting withdrawal: %w", err)
	}
	return nil
}

func (s *partnerService) GetPartnerRewards(partnerID int) ([]models.Reward, error) {
	rewards, err := s.partnerRepo.GetPartnerRewards(partnerID)
	if err != nil {
		return nil, err
	}
	return rewards, nil
}

func (s *partnerService) GetBookings(partnerID int) ([]models.Booking, error) {
	bookings, err := s.bookingRepo.GetBookings(context.Background(), partnerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching bookings for partner %d: %w", partnerID, err)
	}
	return bookings, nil
}
