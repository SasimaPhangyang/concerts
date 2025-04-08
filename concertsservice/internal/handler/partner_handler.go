package handler

import (
	"concerts/internal/models"
	"concerts/internal/repository"
	"concerts/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PartnerHandler struct {
	Service      service.PartnerService
	WithdrawRepo repository.WithdrawRepository
}

func NewPartnerHandler(service service.PartnerService, withdrawRepo repository.WithdrawRepository) *PartnerHandler {
	return &PartnerHandler{
		Service:      service,
		WithdrawRepo: withdrawRepo,
	}
}

func (h *PartnerHandler) GetPartnerBalance(c *gin.Context) {
	partnerID, err := strconv.Atoi(c.Param("partner_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid partner ID"})
		return
	}

	balance, err := h.Service.GetPartnerBalance(partnerID) // ไม่มี ctx แล้ว
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (h *PartnerHandler) GetBookings(c *gin.Context) {
	// ดึง partner_id จาก URL parameters
	partnerIDstr := c.Param("partner_id")
	partnerID, err := strconv.Atoi(partnerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid partner_id"})
		return
	}

	// ดึงข้อมูลการจอง
	bookings, err := h.Service.GetBookings(partnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func (h *PartnerHandler) GetPartnerRewards(c *gin.Context) {
	partnerIDstr := c.Param("partner_id")
	partnerID, err := strconv.Atoi(partnerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid partner_id"})
		return
	}

	rewards, err := h.Service.GetPartnerRewards(partnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch partner rewards"})
		return
	}

	c.JSON(http.StatusOK, rewards)
}

func (h *PartnerHandler) SetAutoWithdraw(c *gin.Context) {
	partnerID, err := strconv.Atoi(c.Param("partner_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid partner ID"})
		return
	}

	var req models.AutoWithdraw
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enabled flag"})
		return
	}

	err = h.Service.SetAutoWithdraw(partnerID, req.Enabled)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Auto-withdraw updated"})
}

// CreateWithdrawRequest สร้างการถอนเงิน
func (h *PartnerHandler) CreateWithdrawRequest(c *gin.Context) {
	partnerIDstr := c.Param("partner_id")
	partnerID, err := strconv.Atoi(partnerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid partner ID"})
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.WithdrawRepo.CreateWithdrawRequest(partnerID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Withdraw request created successfully"})
}

// GetWithdrawRequests ดึงรายการการถอนเงิน
func (h *PartnerHandler) GetWithdrawRequests(c *gin.Context) {
	partnerIDstr := c.Param("partner_id")
	partnerID, err := strconv.Atoi(partnerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid partner ID"})
		return
	}

	withdrawRequests, err := h.WithdrawRepo.GetWithdrawRequests(partnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, withdrawRequests)
}
