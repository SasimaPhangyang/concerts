package handler

import (
	"concerts/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommissionHandler struct {
	Service service.CommissionService
}

func NewCommissionHandler(service service.CommissionService) *CommissionHandler {
	return &CommissionHandler{Service: service}
}

func (h *CommissionHandler) GetCommissions(c *gin.Context) {
	partnerIDstr := c.Param("partner_id")
	partnerID, err := strconv.Atoi(partnerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "partner_id is required"})
		return
	}

	commissions, err := h.Service.GetCommissions(c.Request.Context(), partnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch commissions"})
		return
	}
	c.JSON(http.StatusOK, commissions)
}
