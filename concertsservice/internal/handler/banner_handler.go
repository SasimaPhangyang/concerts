package handler

import (
	"concerts/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	Service service.BannerService // ใช้ interface โดยตรง
}

func NewBannerHandler(service service.BannerService) *BannerHandler {
	return &BannerHandler{Service: service}
}

func (h *BannerHandler) GetBanners(c *gin.Context) {
	banners, err := h.Service.GetBanners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banners"})
		return
	}
	c.JSON(http.StatusOK, banners)
}
