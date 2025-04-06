package handler

import (
	"concerts/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PartnerHandler struct {
	Service service.PartnerService
}

func NewPartnerHandler(service service.PartnerService) *PartnerHandler {
	return &PartnerHandler{Service: service}
}

func (h *PartnerHandler) GetPartnerBalance(c *gin.Context) {
	partnerID := c.DefaultQuery("partner_id", "")
	balance, err := h.Service.GetPartnerBalance(partnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch partner balance"})
		return
	}
	c.JSON(http.StatusOK, balance)
}

// ใน handler/PartnerHandler.go
func (h *PartnerHandler) GetBookings(c *gin.Context) {
	partnerID := c.DefaultQuery("partner_id", "")
	if partnerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partner ID is required"})
		return
	}

	bookings, err := h.Service.GetBookings(partnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func (h *PartnerHandler) GetPartnerRewards(c *gin.Context) {
	partnerID := c.DefaultQuery("partner_id", "") // ดึงค่า partner_id จาก query parameter
	if partnerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partner ID is required"})
		return
	}

	rewards, err := h.Service.GetPartnerRewards(partnerID) // เรียกฟังก์ชันใน PartnerService เพื่อดึงข้อมูลรางวัล
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch partner rewards"})
		return
	}

	c.JSON(http.StatusOK, rewards) // ส่งข้อมูลรางวัลกลับไปใน response
}

func (h *PartnerHandler) SetAutoWithdraw(c *gin.Context) {
	partnerID := c.PostForm("partner_id")
	thresholdStr := c.PostForm("threshold")
	withdrawalMethod := c.PostForm("withdrawal_method")

	// แปลง threshold เป็น float64
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold"})
		return
	}

	err = h.Service.SetAutoWithdraw(partnerID, threshold, withdrawalMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set auto withdrawal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Auto withdrawal set successfully"})
}

func (h *PartnerHandler) RequestWithdrawal(c *gin.Context) {
	partnerID := c.PostForm("partner_id")
	amountStr := c.PostForm("amount")

	// แปลง amount เป็น float64
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	err = h.Service.RequestWithdrawal(partnerID, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request withdrawal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Withdrawal request successful"})
}
