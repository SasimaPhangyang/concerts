package handler

import (
	"concerts/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	Service service.CampaignService
}

func NewCampaignHandler(service service.CampaignService) *CampaignHandler {
	return &CampaignHandler{Service: service}
}

func (h *CampaignHandler) JoinCampaign(c *gin.Context) {
	partnerID := c.PostForm("partner_id")
	campaignID := c.PostForm("campaign_id")

	err := h.Service.JoinCampaign(partnerID, campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join campaign"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
